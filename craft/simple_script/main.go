package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/l-f-h/atcompiler/craft/parser"
)

func main() {
	verbose := flag.Bool("verbose", false, "打印语法树")
	flag.Parse()
	simpleParser := parser.NewSimpleParser()
	prompt := "\n>"
	simpleScript := NewSimpleScript(*verbose)
	scriptText := ""
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Printf(prompt)
		scanner.Scan()
		line := scanner.Text()
		scriptText += line + "\n"
		if strings.HasSuffix(line, ";") {
			node, err := simpleParser.Parse(scriptText)
			if err != nil {
				fmt.Println("parse error: ", err)
				continue
			}
			if *verbose {
				parser.DumpAST(node, "")
			}
			simpleScript.Evaluate(node, "")
			scriptText = ""
		}
	}

}
