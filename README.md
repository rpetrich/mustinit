# mustinit

A simple go analysis pass that reports on uninitialized variables and fields.
Initialization requirements are specified via mustinit comments or struct tags
and automatically inferred from common golang patterns.

```bash
$ go get github.com/rpetrich/mustinit/cmd/mustinit
$ go vet -vettool=`which mustinit` ./samples
# github.com/rpetrich/mustinit/samples
samples/all.go:19:3: bare return of type literal requiring initialization
samples/all.go:21:9: missing required fields in literal literal: required
samples/all.go:27:9: literal fields must be specified by key since literal has required fields: required
samples/all.go:31:6: type literal requires initialization
samples/all.go:40:9: missing required fields in embedded literal: lit
samples/all.go:46:6: type initializedInt requires initialization
samples/all.go:51:6: type io.PipeWriter requires initialization
samples/all.go:56:9: missing required fields in mustinit.TypeRequirements literal: IsRequired, RequiredFields
```