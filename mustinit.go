package mustinit

import (
	"flag"
	"fmt"
	"go/ast"
	"go/token"
	"go/types"
	"reflect"
	"sort"
	"strings"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

var Analyzer = &analysis.Analyzer{
	Name:      "mustinit",
	Doc:       "checks for initialization of required fields",
	Flags:     *flag.NewFlagSet("", flag.ExitOnError),
	Requires:  []*analysis.Analyzer{inspect.Analyzer},
	Run:       run,
	FactTypes: []analysis.Fact{new(PackageFact)},
}

func init() {
	Analyzer.Flags.Var(&globalDefaultRequirement, "default-init-requirement", "default initialization requirement to apply to sources (one of: none, values, fields, all)")
}

type initRequirement int // mustinit:true

const (
	initRequirementNone   initRequirement = 0
	initRequirementValues                 = 1 << iota
	initRequirementFields
	initRequirementSkip
)

func (ir *initRequirement) String() string {
	if ir == nil {
		return "none"
	}
	switch *ir {
	case initRequirementNone:
		return "none"
	case initRequirementValues:
		return "values"
	case initRequirementFields:
		return "fields"
	case (initRequirementValues | initRequirementFields):
		return "all"
	case initRequirementSkip:
		return "skip"
	default:
		return "unknown"
	}
}

func (ir *initRequirement) Set(value string) error {
	switch value {
	case "none":
		*ir = initRequirementNone
		return nil
	case "values":
		*ir = initRequirementValues
		return nil
	case "fields":
		*ir = initRequirementFields
		return nil
	case "all":
		*ir = (initRequirementValues | initRequirementFields)
		return nil
	case "skip":
		*ir = initRequirementSkip
		return nil
	default:
		return fmt.Errorf("invalid value %q", value)
	}
}

var globalDefaultRequirement initRequirement

// TypeRequiremends describe whether a type and its subfields must be initialized
type TypeRequirements struct {
	// IsRequired indicates the type must be initialized
	IsRequired bool `mustinit:"true"`
	// RequiredFIelds indicates which subfields must be initialized
	RequiredFields map[string]struct{} `mustinit:"true"`
}

// PackageFact is used to transfer data on packages between passes
type PackageFact struct {
	// Requirements stores the requirements of all types in the package
	Requirements map[string]TypeRequirements `mustinit:"true"`
}

// AFact is a type tag required to conform to analysis.Fact
func (*PackageFact) AFact() {}

type declarationSource struct {
	typeSpec *ast.TypeSpec  `mustinit:"true"`
	comments []*ast.Comment `mustinit:"true"`
}

// factStore is used to incrmentally analyze initialization requirements
type factStore struct {
	ourFacts            PackageFact                                       `mustinit:"true"`
	pendingDeclarations map[string]declarationSource                      `mustinit:"true"`
	otherFacts          map[string]PackageFact                            `mustinit:"true"`
	packages            map[string]*types.Package                         `mustinit:"true"`
	otherFactImporter   func(pkg *types.Package, fact analysis.Fact) bool `mustinit:"true"`
	defaultRequirement  initRequirement                                   `mustinit:"true"`
}

// newFactStore creates a fact store to perform an analysis pass
func newFactStore(pass *analysis.Pass) factStore {
	requirements := map[string]TypeRequirements{}
	defaultRequirement := globalDefaultRequirement
	if packageDefault, ok := defaultRequirements[pass.Pkg.Path()]; ok {
		defaultRequirement = packageDefault.requirement
		if packageDefault.types != nil {
			for k, v := range packageDefault.types {
				requirements[k] = v
			}
		}
	}
	return factStore{
		ourFacts: PackageFact{
			Requirements: requirements,
		},
		pendingDeclarations: map[string]declarationSource{},
		otherFacts:          map[string]PackageFact{},
		packages:            map[string]*types.Package{},
		otherFactImporter:   pass.ImportPackageFact,
		defaultRequirement:  defaultRequirement,
	}
}

// analyzeRequirements parses struct tags to determine which fields require
// initialization
func (facts *factStore) analyzeRequirements(decl declarationSource) TypeRequirements {
	result := TypeRequirements{
		RequiredFields: nil,
		IsRequired:     false,
	}
	requirement := facts.defaultRequirement
	for _, comment := range decl.comments {
		switch comment.Text {
		case "// mustinit:true":
			requirement = initRequirementValues
		case "// mustinit:false":
			requirement = initRequirementNone
		case "// mustinit:values":
			requirement = initRequirementValues
		case "// mustinit:fields":
			requirement = initRequirementFields
		case "// mustinit:all":
			requirement = initRequirementValues | initRequirementFields
		case "// mustinit:skip":
			return result
		}
	}
	if requirement&initRequirementValues == initRequirementValues {
		result.IsRequired = true
	}
	switch typeDef := decl.typeSpec.Type.(type) {
	case *ast.StructType:
		requiredFields := map[string]struct{}{}
		for _, field := range typeDef.Fields.List {
			// check if field has a mustinit:"true" tag
			found := false
			if field.Tag != nil && field.Tag.Kind == token.STRING {
				tag := reflect.StructTag(parseStringLiteral(field.Tag.Value))
				var value bool
				if value, found = mustInitFromTag(tag); found && value {
					for _, name := range field.Names {
						requiredFields[name.Name] = struct{}{}
					}
					continue
				}
			}
			// check if there's a default requirement that fields are required
			if !found && requirement&initRequirementFields == initRequirementFields {
				for _, name := range field.Names {
					if name.Name != "_" {
						requiredFields[name.Name] = struct{}{}
					}
				}
				continue
			}
			// check if field's type requires initialization
			if field.Type != nil {
				if requirements, ok := facts.Lookup(field.Type); ok && requirements.IsRequired {
					for _, name := range field.Names {
						if name.Name != "_" {
							requiredFields[name.Name] = struct{}{}
						}
					}
					continue
				}
			}
		}
		result.RequiredFields = requiredFields
		if !result.IsRequired {
			result.IsRequired = len(requiredFields) != 0
		}
	default:
		result, _ = facts.Lookup(typeDef)
	}
	return result
}

// AddTypeSpec loads a type spec for future analysis
func (facts *factStore) AddTypeSpec(n *ast.TypeSpec, comments []*ast.Comment) {
	facts.pendingDeclarations[n.Name.Name] = declarationSource{
		typeSpec: n,
		comments: comments,
	}
}

// AddImportSpec loads an import spec and associates its identifier with the
// appropriate package
func (facts *factStore) AddImportSpec(info *types.Info, importSpec *ast.ImportSpec) {
	obj, ok := info.Implicits[importSpec]
	if !ok {
		obj = info.Defs[importSpec.Name] // renaming import
	}
	pkg := obj.(*types.PkgName).Imported()
	if ok {
		facts.packages[pkg.Name()] = pkg
	} else {
		facts.packages[importSpec.Name.Name] = pkg
	}
}

// Lookup looks up a type reference for its type requriements, performing
// analysis if necessary
func (facts *factStore) Lookup(typeRef ast.Expr) (TypeRequirements, bool) {
	switch typeRef := typeRef.(type) {
	case *ast.Ident:
		// lookup in our package
		if result, ok := facts.ourFacts.Requirements[typeRef.Name]; ok {
			return result, true
		}
		// lazily populate within our package
		if decl, ok := facts.pendingDeclarations[typeRef.Name]; ok {
			result := facts.analyzeRequirements(decl)
			facts.ourFacts.Requirements[typeRef.Name] = result
			return result, true
		}
	case *ast.SelectorExpr:
		if x, ok := typeRef.X.(*ast.Ident); ok {
			other, ok := facts.otherFacts[x.Name]
			if ok {
				result, ok := other.Requirements[typeRef.Sel.Name]
				return result, ok
			}
			if pkg, ok := facts.packages[x.Name]; ok {
				if facts.otherFactImporter(pkg, &other) {
					facts.otherFacts[x.Name] = other
					result, ok := other.Requirements[typeRef.Sel.Name]
					return result, ok
				}
			}
		}
	}
	// unknown, assume no requirements
	return TypeRequirements{
		IsRequired:     false,
		RequiredFields: nil,
	}, false
}

func (facts *factStore) ShouldSkip() bool {
	return facts.defaultRequirement&initRequirementSkip == initRequirementSkip
}

// Analyze performs full analysis so that facts about this package may be
// imported by other packages in their analysis
func (facts *factStore) Analyze() analysis.Fact {
	for k, v := range facts.pendingDeclarations {
		if _, ok := facts.ourFacts.Requirements[k]; !ok {
			facts.ourFacts.Requirements[k] = facts.analyzeRequirements(v)
		}
	}
	return &facts.ourFacts
}

func run(pass *analysis.Pass) (interface{}, error) {
	inspect := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)

	// collect types and imports
	facts := newFactStore(pass)
	commentMaps := make([]ast.CommentMap, len(pass.Files))
	for i, file := range pass.Files {
		commentMaps[i] = ast.NewCommentMap(pass.Fset, file, file.Comments)
	}
	inspect.WithStack(nil, func(n ast.Node, push bool, stack []ast.Node) bool {
		if push {
			switch n := n.(type) {
			case *ast.ImportSpec:
				facts.AddImportSpec(pass.TypesInfo, n)
			case *ast.TypeSpec:
				facts.AddTypeSpec(n, findAllComments(commentMaps, stack))
			}
		}
		return true
	})

	// inspect all uses of types
	if !facts.ShouldSkip() {
		inspect.WithStack(nil, func(n ast.Node, push bool, stack []ast.Node) bool {
			if !push {
				return true
			}
			switch n := n.(type) {
			case *ast.CompositeLit:
				if reqs, found := facts.Lookup(n.Type); found && len(reqs.RequiredFields) != 0 {
					foundFields := map[string]struct{}{}
					for _, v := range n.Elts {
						if pair, ok := v.(*ast.KeyValueExpr); ok {
							if key, ok := pair.Key.(*ast.Ident); ok {
								if _, found := reqs.RequiredFields[key.Name]; found {
									foundFields[key.Name] = struct{}{}
								}
								continue
							}
						}
						pass.Reportf(n.Pos(), "fields must be initialized by key when any fields are required")
						return true
					}
					if len(foundFields) != len(reqs.RequiredFields) {
						missingFields := map[string]struct{}{}
						for field := range reqs.RequiredFields {
							if _, ok := foundFields[field]; !ok {
								missingFields[field] = struct{}{}
							}
						}
						pass.Reportf(n.Pos(), "missing required fields in literal: %s", stringifyRequiredFields(missingFields))
					}
				}
			case *ast.ValueSpec:
				if len(n.Values) == 0 {
					if reqs, found := facts.Lookup(n.Type); found && reqs.IsRequired {
						pass.Reportf(n.Pos(), "value type requires initialization")
					}
				}
			case *ast.ReturnStmt:
				if len(n.Results) == 0 {
					if fn, ok := findFunction(stack); ok && fn.Type.Results != nil && len(fn.Type.Results.List) != 0 {
						for _, field := range fn.Type.Results.List {
							if reqs, ok := facts.Lookup(field.Type); ok && reqs.IsRequired {
								pass.Reportf(n.Pos(), "bare return of type requiring initialization")
								return true
							}
						}
					}
				}
			}
			return true
		})
	}

	// export a fact dsecribing this package
	pass.ExportPackageFact(facts.Analyze())

	return nil, nil
}

func parseStringLiteral(lit string) string {
	if len(lit) >= 2 {
		if lit[0] == '`' {
			return lit[1 : len(lit)-1]
		}
		if lit[0] == '"' {
			// TODO: apply character escapes
			return lit[1 : len(lit)-1]
		}
	}
	return ""
}

func mustInitFromTag(tag reflect.StructTag) (result bool, found bool) {
	if value, ok := tag.Lookup("mustinit"); ok {
		return value == "true", true
	}
	return false, false
}

func stringifyRequiredFields(fieldNames map[string]struct{}) string {
	fields := make([]string, 0, len(fieldNames))
	for field := range fieldNames {
		fields = append(fields, field)
	}
	sort.Strings(fields)
	return strings.Join(fields, ", ")
}

func findFunction(stack []ast.Node) (*ast.FuncDecl, bool) {
	for i := len(stack) - 1; i >= 0; i-- {
		if result, ok := stack[i].(*ast.FuncDecl); ok {
			return result, true
		}
	}
	return nil, false
}

func findAllComments(commentMaps []ast.CommentMap, stack []ast.Node) []*ast.Comment {
	var result []*ast.Comment
	for i := len(stack) - 1; i >= 0; i-- {
		node := stack[i]
		for _, commentMap := range commentMaps {
			if comments, ok := commentMap[node]; ok {
				if result == nil && len(comments) == 1 {
					result = comments[0].List
				} else {
					for _, commentGroup := range comments {
						result = append(result, commentGroup.List...)
					}
				}
			}
		}
	}
	return result
}
