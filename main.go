package main

import (
	"bufio"
	"os"

	"github.com/wreulicke/gojg/generator"
	"github.com/wreulicke/gojg/parse"
)

func main() {
	ast, e := parse.Parse("true")
	if e != nil {
		return
	}
	writer := bufio.NewWriter(os.Stdout)
	err := generator.Generate(ast, writer)
	if err != nil {
		return
	}
	writer.Flush()
}
