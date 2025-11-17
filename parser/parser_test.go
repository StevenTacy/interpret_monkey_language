package parser

import (
	"interpreter/ast"
	"interpreter/lexer"
	"testing"
)

func TestLetStatement(t *testing.T) {
	input := `
let x = 5;
let y = 10;
let foobar = 838383;
`
	l := lexer.New(input)
	p := New(l)
	program := p.ParseProgram()
	checkParserErrors(t, p)

	if program == nil {
		t.Fatalf("ParseProgram return nil")
	}

	if len(program.Statements) != 3 {
		t.Fatalf("program.Statements malfunctioned expected %d, but got %d", 3, len(program.Statements))
	}

	tests := []struct {
		expectedIdentifier string
	}{
		{"x"},
		{"y"},
		{"foobar"},
	}

	for i, test := range tests {
		stmt := program.Statements[i]
		if !testLetStatement(t, stmt, test.expectedIdentifier) {
			return
		}
	}
}

func testLetStatement(t *testing.T, s ast.Statement, want string) bool {
	if s.TokenLiteral() != "let" {
		t.Errorf("s.TokenLiteral not has expected value %s, got %s", "let", s.TokenLiteral())
	}

	letStmt, ok := s.(*ast.LetStatement)
	if !ok {
		t.Errorf("s is not a letStatement, got %T", s)
		return false
	}

	if letStmt.Name.Value != want {
		t.Errorf("the Name field has differences got %s, but want %s", letStmt.Name, want)
		return false
	}

	if letStmt.Name.TokenLiteral() != want {
		t.Errorf("s.Name has wrong literal got %s, but want %s", letStmt.Name, want)
		return false
	}

	return true
}

func checkParserErrors(t *testing.T, p *Parser) {
	errors := p.errors
	if len(errors) == 0 {
		return
	}

	t.Errorf("parser has %d errors", len(errors))
	for _, msg := range errors {
		t.Errorf("parser error: %q", msg)
	}
	t.FailNow()
}
