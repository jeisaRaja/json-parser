package lexer

import (
	"jeisaraja/json_parser/token"
	"testing"
)

func TestLexer(t *testing.T) {
	tests := []struct {
		input          string
		expectedTokens []struct {
			expectedType    token.TokenType
			expectedLiteral string
		}
	}{
		{
			`{"yeah": "name"}`,
			[]struct {
				expectedType    token.TokenType
				expectedLiteral string
			}{
				{token.LBRACE, "{"},
				{token.STRING, "yeah"},
				{token.COLON, ":"},
				{token.STRING, "name"},
				{token.RBRACE, "}"},
				{token.EOF, ""},
			},
		},
		{
			`{"num": 123, "bool": true, "nullValue": null}`,
			[]struct {
				expectedType    token.TokenType
				expectedLiteral string
			}{
				{token.LBRACE, "{"},
				{token.STRING, "num"},
				{token.COLON, ":"},
				{token.NUMBER, "123"},
				{token.COMMA, ","},
				{token.STRING, "bool"},
				{token.COLON, ":"},
				{token.TRUE, "true"},
				{token.COMMA, ","},
				{token.STRING, "nullValue"},
				{token.COLON, ":"},
				{token.NULL, "null"},
				{token.RBRACE, "}"},
				{token.EOF, ""},
			},
		},
		{
			`[{"key": "value"}, {"arr": [1, 2, 3]}]`,
			[]struct {
				expectedType    token.TokenType
				expectedLiteral string
			}{
				{token.LBRACKET, "["},
				{token.LBRACE, "{"},
				{token.STRING, "key"},
				{token.COLON, ":"},
				{token.STRING, "value"},
				{token.RBRACE, "}"},
				{token.COMMA, ","},
				{token.LBRACE, "{"},
				{token.STRING, "arr"},
				{token.COLON, ":"},
				{token.LBRACKET, "["},
				{token.NUMBER, "1"},
				{token.COMMA, ","},
				{token.NUMBER, "2"},
				{token.COMMA, ","},
				{token.NUMBER, "3"},
				{token.RBRACKET, "]"},
				{token.RBRACE, "}"},
				{token.RBRACKET, "]"},
				{token.EOF, ""},
			},
		},
		{
			`{"stringWithEscape": "Hello\nWorld"}`,
			[]struct {
				expectedType    token.TokenType
				expectedLiteral string
			}{
				{token.LBRACE, "{"},
				{token.STRING, "stringWithEscape"},
				{token.COLON, ":"},
				{token.STRING, "Hello\nWorld"},
				{token.RBRACE, "}"},
				{token.EOF, ""},
			},
		},
	}

	for _, tt := range tests {
		l := New(tt.input)

		for i, et := range tt.expectedTokens {
			tok := l.NextToken()
			t.Log(tt.input, tok)
			if tok.Type != et.expectedType {
				t.Fatalf("test case %q - tests[%d] - token type wrong. expected %q, got %q", tt.input, i, et.expectedType, tok.Type)
			}
			if tok.Literal != et.expectedLiteral {
				t.Fatalf("test case %q - tests[%d] - token literal wrong. expected %q, got %q", tt.input, i, et.expectedLiteral, tok.Literal)
			}
		}
	}
}

