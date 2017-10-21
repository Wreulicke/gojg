package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/alecthomas/kingpin"
	"github.com/wreulicke/gojg/generator"
	"github.com/wreulicke/gojg/parse"
)

var (
	verbose    = kingpin.Flag("verbose", "Set verbose mode").Short('v').Bool()
	template   = kingpin.Arg("template", "Template File").Required().String()
	contextMap = kingpin.Flag("context", "values").Short('c').StringMap()
)

func main() {
	kingpin.Version("0.0.1")
	kingpin.CommandLine.HelpFlag.Short('h')
	kingpin.Parse()

	f, e := os.Open(*template)
	if e != nil {
		fmt.Println(e)
		os.Exit(1)
	}

	reader := bufio.NewReader(f)
	ast, e := parse.Parse(reader)
	if e != nil {
		f.Close()
		fmt.Println(e)
		os.Exit(1)
	}

	writer := bufio.NewWriter(os.Stdout)

	context := make(map[string]interface{}, len(*contextMap))
	for k, v := range *contextMap {
		context[k] = v
	}

	g := generator.NewGenerator(context, writer)
	e = g.Generate(ast)
	if e != nil {
		f.Close()
		fmt.Println(e)
		os.Exit(1)
	}

	writer.WriteRune('\n')
	writer.Flush()
}
