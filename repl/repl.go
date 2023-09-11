package repl

import (
	"bufio"
	"fmt"
	"io"
	"kulang/lexer"
	"kulang/token"
)

const PROMPT = "> "

func Start(in io.Reader, out io.Writer) {
	sc := bufio.NewScanner(in)

	for {
		fmt.Fprintf(out, PROMPT)

		scanned := sc.Scan()
		if !scanned {
			return
		}

		line := sc.Text()
		if line == "exit" || line == "q" || line == "quit" {
			return
		}

		l := lexer.New(line)
		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Fprintf(out, "%+v\n", tok)
		}
	}
}
