package lexer

import (
	"monkey/token"
	"testing"
)

func TestNextToken(t *testing.T) {
	input := `=+(){},;`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.ASSIGN, "="},
		{token.PLUS, "+"},
		{token.LPAREN, "("},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RBRACE, "}"},
		{token.COMMA, ","},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
	}

	l := New(input)

	for i, tokenType := range tests {
		tok := l.NextToken()

		if tok.Type != tokenType.expectedType {
			t.Fatalf("tests[%d] - tokenType wrong. Expected=%q, got=%q",
				i, tokenType.expectedType, tok.Type)
		}
		if tok.Literal != tokenType.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. Expected=%q, got=%q",
				i, tokenType.expectedLiteral, tok.Literal)
		}
	}
}
