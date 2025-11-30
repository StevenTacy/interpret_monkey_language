package repl

import (
	"bufio"
	"fmt"
	"interpreter/evaluator"
	"interpreter/lexer"
	"interpreter/object"
	"interpreter/parser"
	"io"
)

const PROMPT = ">> "

const MONKEY_FACE = `       __,__
.--. .-"
 "-. .--.
 / .. \/ .-. .-. \/ .. \
 | | '| / Y \ |' ||
 | \ \ \ 0 | 0 / / /|
 \ '- ,\.-"""""""-./,-' /
 ''-' /_ ^ ^ _\ '-''
	 | \._ _./ |
	 \ \ '~' / /
	 '._ '-=-' _.'
		 '-----'
`

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	env := object.NewEnvironment()

	for {
		fmt.Printf(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}
		line := scanner.Text()
		l := lexer.New(line)
		p := parser.New(l)

		program := p.ParseProgram()
		if len(p.Errors()) != 0 {
			printParseErrors(out, p.Errors())
			continue
		}

		evaluated := evaluator.Eval(program, env)
		if evaluated != nil {
			io.WriteString(out, evaluated.Inspect())
			io.WriteString(out, "\n")
		}
	}
}

func printParseErrors(w io.Writer, errs []string) {
	io.WriteString(w, MONKEY_FACE)
	io.WriteString(w, "Oops, we got some issues here\n")
	io.WriteString(w, "Parser errors: \n")
	for _, err := range errs {
		io.WriteString(w, "\t"+err+"\n")
	}
}
