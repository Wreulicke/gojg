package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/alecthomas/kingpin"
	"github.com/wreulicke/gojg/generator"
	"github.com/wreulicke/gojg/parse"
)

func resolveTemplate() *bufio.Reader {
	if *template != nil {
		return bufio.NewReader(*template)
	}
	return bufio.NewReader(os.Stdin)
}

func close() {
	if *template != nil {
		(*template).Close()
	}
}

var (
	verbose    = kingpin.Flag("verbose", "Set verbose mode").Short('v').Bool()
	contextMap = kingpin.Flag("context", "Context Parameter").Short('c').StringMap()
	template   = kingpin.Arg("template", "Template File").File()
)

func main() {
	kingpin.Version("0.0.1")
	kingpin.CommandLine.HelpFlag.Short('h')
	kingpin.Parse()

	reader := resolveTemplate()

	ast, e := parse.Parse(reader)
	if e != nil {
		fmt.Println(e)
		close()
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
		close()
		fmt.Println(e)
		os.Exit(1)
	}

	writer.WriteRune('\n')
	writer.Flush()
}
