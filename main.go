package main

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/chzyer/readline"
	"io"
	"github.com/sheik/antlr4-go-visitor-test/parser"
)

func main() {
	l, err := readline.NewEx(&readline.Config{
		Prompt:          "calc> ",
		HistoryFile:     "/tmp/readline.tmp",
		InterruptPrompt: "^C",
		EOFPrompt:       "exit",
		HistorySearchFold:   true,
	})
	if err != nil {
		panic(err)
	}
	defer l.Close()

	baseParser := new(parser.BaseExprVisitor)

	for {
		line, err := l.Readline()
		if err == readline.ErrInterrupt {
			if len(line) == 0 {
				break
			} else {
				continue
			}
		} else if err == io.EOF {
			break
		}
		input := antlr.NewInputStream(line+"\r\n")
		lexer := parser.NewExprLexer(input)
		stream := antlr.NewCommonTokenStream(lexer,0)
		p := parser.NewExprParser(stream)
		p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
		p.BuildParseTrees = true
		tree := p.Prog()
		baseParser.Visit(tree)
	}
}
