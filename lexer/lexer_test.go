package lexer

import (
	"monkeylang/token"
	"testing"
)

func TestNextToken(t *testing.T) {
	input := []string{
		"=+(){},;",
		`let five = 5;
		let ten = 10;
		let add = fn(x, y) {
		x + y;
		};
		let result = add(five, ten);`,
	}

	expected := [][]token.Token{
		{
			{token.ASSIGN, "="},
			{token.PLUS, "+"},
			{token.LPAREN, "("},
			{token.RPAREN, ")"},
			{token.LBRACE, "{"},
			{token.RBRACE, "}"},
			{token.COMMA, ","},
			{token.SEMICOLON, ";"},
			{token.EOF, ""},
		}, {
			{token.LET, "let"},
			{token.IDENT, "five"},
			{token.ASSIGN, "="},
			{token.INT, "5"},
			{token.SEMICOLON, ";"},
			{token.LET, "let"},
			{token.IDENT, "ten"},
			{token.ASSIGN, "="},
			{token.INT, "10"},
			{token.SEMICOLON, ";"},
			{token.LET, "let"},
			{token.IDENT, "add"},
			{token.ASSIGN, "="},
			{token.FUNCTION, "fn"},
			{token.LPAREN, "("},
			{token.IDENT, "x"},
			{token.COMMA, ","},
			{token.IDENT, "y"},
			{token.RPAREN, ")"},
			{token.LBRACE, "{"},
			{token.IDENT, "x"},
			{token.PLUS, "+"},
			{token.IDENT, "y"},
			{token.SEMICOLON, ";"},
			{token.RBRACE, "}"},
			{token.SEMICOLON, ";"},
			{token.LET, "let"},
			{token.IDENT, "result"},
			{token.ASSIGN, "="},
			{token.IDENT, "add"},
			{token.LPAREN, "("},
			{token.IDENT, "five"},
			{token.COMMA, ","},
			{token.IDENT, "ten"},
			{token.RPAREN, ")"},
			{token.SEMICOLON, ";"},
			{token.EOF, ""},
		},
	}
	for testIdx := range input {
		l := New(input[testIdx])
		for i, tt := range expected[testIdx] {
			tok := l.NextToken()

			if tok.Type != tt.Type {
				t.Fatalf("tests[%d][%d] - tokentype wrong. expected=%q, got=%q",
					testIdx, i, tt.Type, tok.Type)
			}

			if tok.Literal != tt.Literal {
				t.Fatalf("tests[%d][%d] - literal wrong. expected=%q, got=%q",
					testIdx, i, tt.Literal, tok.Literal)
			}
		}
	}
}