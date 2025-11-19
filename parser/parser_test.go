package parser

import (
	"fmt"
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
	// input := `
	// let x 5;
	// let = 10;
	// let 838383;
	// `
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

func TestReturnStatements(t *testing.T) {
	input := `
 return 5;
 return 10;
 return 993322;
 `
	l := lexer.New(input)
	p := New(l)
	program := p.ParseProgram()
	checkParserErrors(t, p)

	if len(program.Statements) != 3 {
		t.Errorf("program.Statements errored want %d, but got %d", 3, len(program.Statements))
	}

	for _, stmt := range program.Statements {
		returnStmt, ok := stmt.(*ast.ReturnStatement)
		if !ok {
			t.Errorf("stmt is not *ast.ReturnStatement got %T want %T", stmt, returnStmt)
			continue
		}

		if returnStmt.TokenLiteral() != "return" {
			t.Errorf("returnStmt has wrong literal got %q, but want %q", returnStmt.TokenLiteral(), "return")
		}
	}
}

func TestIdentifierExpression(t *testing.T) {
	input := "foobar;"
	l := lexer.New(input)
	p := New(l)
	program := p.ParseProgram()
	checkParserErrors(t, p)

	if len(program.Statements) != 1 {
		t.Fatalf("program do not  have enough statement got %d want %d", len(program.Statements), 1)
	}

	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("program.Statements[0] is not ast.ExpressionStatement, got %T", program.Statements[0])
	}

	ident, ok := stmt.Expression.(*ast.Identifier)
	if !ok {
		t.Fatalf("exp is not ast.Identifier, got %T", stmt.Expression)
	}

	if ident.Value != "foobar" {
		t.Errorf("idtentifier value is wrong got %s want %s", ident.Value, "foobar")
	}

	if ident.TokenLiteral() != "foobar" {
		t.Errorf("idtentifier literal is wrong got %s want %s", ident.TokenLiteral(), "foobar")
	}
}

func TestIntegerExpression(t *testing.T) {
	input := "5;"
	l := lexer.New(input)
	p := New(l)
	program := p.ParseProgram()
	checkParserErrors(t, p)

	if len(program.Statements) != 1 {
		t.Fatalf("program do not  have enough statement got %d want %d", len(program.Statements), 1)
	}

	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("program.Statements[0] is not ast.ExpressionStatement, got %T", program.Statements[0])
	}

	ident, ok := stmt.Expression.(*ast.IntegerLiteral)
	if !ok {
		t.Fatalf("exp is not ast.IntegerLiteral, got %T", stmt.Expression)
	}

	if ident.Value != 5 {
		t.Errorf("Integer value is wrong got %d want %d", ident.Value, 5)
	}

	if ident.TokenLiteral() != "5" {
		t.Errorf("Integer literal is wrong got %s want %s", ident.TokenLiteral(), "5")
	}
}

func TestPrefixExpression(t *testing.T) {
	prefixTests := []struct {
		input    string
		operator string
		value    int64
	}{
		{"!5;", "!", 5},
		{"-15;", "-", 15},
	}

	for _, test := range prefixTests {
		l := lexer.New(test.input)
		p := New(l)
		program := p.ParseProgram()
		checkParserErrors(t, p)

		if len(program.Statements) != 1 {
			t.Fatalf("program do not  have enough statement got %d want %d", len(program.Statements), 1)
		}

		stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
		if !ok {
			t.Fatalf("program.Statements[0] is not ast.ExpressionStatement, got %T", program.Statements[0])
		}

		exp, ok := stmt.Expression.(*ast.PrefixExpression)
		if !ok {
			t.Fatalf("exp is not ast.PrefixExpression, got %T", stmt.Expression)
		}

		if exp.Operator != test.operator {
			t.Errorf("Integer value is wrong got %s want %s", exp.Operator, test.operator)
		}

		if !testIntegerLiteral(t, exp.Right, test.value) {
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

func testIntegerLiteral(t *testing.T, il ast.Expression, value int64) bool {
	integ, ok := il.(*ast.IntegerLiteral)
	if !ok {
		t.Errorf("il is not an IntegerLiteral, got %T", il)
		return false
	}

	if integ.Value != value {
		t.Errorf("integ.Value is not i wanted got %d, want %d", integ.Value, value)
		return false
	}

	literal := fmt.Sprintf("%d", value)
	if integ.TokenLiteral() != literal {
		t.Errorf("integ.TokenLiteral() is not i wanted got %s, want %s", integ.TokenLiteral(), literal)
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
