package samples

import (
	"github.com/rpetrich/mustinit"
	"io"
	"strconv"
)

type literal struct {
	required int `mustinit:"true"`
	optional int
}

func returnLiteral() (lit literal, err error) {
	var result int
	result, err = strconv.Atoi("-42")
	if err != nil {
		return
	}
	return literal{
		optional: result,
	}, nil
}

func unnamedLiteral() literal {
	return literal{0, 0}
}

func varDeclaration() literal {
	var result literal
	return result
}

type embedded struct {
	lit literal
}

func returnEmbedded() embedded {
	return embedded{}
}

type initializedInt int // mustinit:true

func returnInt() initializedInt {
	var result initializedInt
	return result
}

func pipeWriter() io.PipeWriter {
	var result io.PipeWriter
	return result
}

func mustInit() mustinit.TypeRequirements {
	return mustinit.TypeRequirements{}
}
