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
	version     string
	hash        string
	buildtime   string
	goversion   string
	verbose     = kingpin.Flag("verbose", "Set verbose mode").Bool()
	contextMap  = kingpin.Flag("context", "Context Parameter").Short('c').StringMap()
	contextFile = kingpin.Flag("context-file", "Context File").Short('f').File()
	details     = kingpin.Flag("details", "Binary Infomation").Short('d').Bool()
	template    = kingpin.Arg("template", "Template File").File()
	outputFile  = kingpin.Flag("output", "Output File").Short('o').
			OpenFile(os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
)

func main() {
	kingpin.Version(version)
	kingpin.CommandLine.VersionFlag.Short('v')
	kingpin.CommandLine.HelpFlag.Short('h')
	kingpin.Parse()

	if *details == true {
		fmt.Printf("version: %s\r\n", version)
		fmt.Printf("commit hash: %s\r\n", hash)
		fmt.Printf("build timestamp: %s\r\n", buildtime)
		fmt.Printf("build go version: %s\r\n", goversion)
		return
	}

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
