package main

import (
	"golang.org/x/tools/go/analysis/singlechecker"

	"github.com/rpetrich/mustinit"
)

func main() {
	singlechecker.Main(mustinit.Analyzer)
}
