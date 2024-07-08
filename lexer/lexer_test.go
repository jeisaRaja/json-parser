package lexer

import (
	"jeisaraja/json_parser/token"
	"testing"
)

func TestLexer(t *testing.T) {
	input := "{}"
	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.LBRACE, "{"},
		{token.RBRACE, "}"},
    {token.EOF, ""},
	}

	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()
		t.Log(tok)
		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - token type wrong. expected %q, got %q", i, tt.expectedType, tok.Type)
		}
		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - token literal wrong. expected %q, got %q", i, tt.expectedLiteral, tok.Type)
		}
	}
}
