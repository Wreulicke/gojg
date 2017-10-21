package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/wreulicke/gojg/context"

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

func output() *bufio.Writer {
	if *outputFile != nil {
		return bufio.NewWriter(*outputFile)
	}
	return bufio.NewWriter(os.Stdout)
}

func close() {
	if *template != nil {
		(*template).Close()
	}
}

var (
	verbose     = kingpin.Flag("verbose", "Set verbose mode").Short('v').Bool()
	contextMap  = kingpin.Flag("context", "Context Parameter").Short('c').StringMap()
	contextFile = kingpin.Flag("context-file", "Context File").Short('f').File()
	template    = kingpin.Arg("template", "Template File").File()
	outputFile  = kingpin.Flag("output", "Output File").Short('o').
			OpenFile(os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
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

	writer := output()

	context, e := context.CreateContext(contextMap, contextFile)
	if e != nil {
		close()
		fmt.Println(e)
		os.Exit(1)
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
