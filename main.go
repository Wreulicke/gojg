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
	verbose  = kingpin.Flag("verbose", "Set verbose mode").Short('v').Bool()
	template = kingpin.Arg("template", "Template File").Required().String()
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
		fmt.Println("error1")
		os.Exit(1)
	}

	writer := bufio.NewWriter(os.Stdout)
	err := generator.Generate(ast, writer)
	writer.WriteRune('\n')
	writer.Flush()
	if err != nil {
		f.Close()
		fmt.Println(e)
		os.Exit(1)
	}

	writer.WriteRune('\n')
	writer.Flush()
}
