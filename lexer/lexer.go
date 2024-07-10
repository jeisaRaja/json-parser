package lexer

import (
	"jeisaraja/json_parser/token"
	"strings"
)

type Lexer struct {
	input        string
	position     int
	readPosition int
	ch           byte
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhitespace()

	switch l.ch {
	case '{':
		tok.Type = token.LBRACE
		tok.Literal = string(l.ch)
	case '}':
		tok.Type = token.RBRACE
		tok.Literal = string(l.ch)
	case ':':
		tok.Type = token.COLON
		tok.Literal = ":"
	case '"':
		tok.Type = token.STRING
		tok.Literal = l.readString()
	case ',':
		tok.Type = token.COMMA
		tok.Literal = ","
	case '[':
		tok.Type = token.LBRACKET
		tok.Literal = "["
	case ']':
		tok.Type = token.RBRACKET
		tok.Literal = "]"
	case 0:
		tok.Type = token.EOF
		tok.Literal = ""
	default:
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = lookupIdent(tok.Literal)
			return tok
		} else if isDigit(l.ch) {
			tok.Type = token.NUMBER
			tok.Literal = l.readNumber()
			return tok
		} else {
			tok.Type = token.ILLEGAL
			tok.Literal = ""
		}
	}
	l.readChar()
	return tok
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition++
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

func isDigit(char byte) bool {
	return char >= '0' && char <= '9'
}

func (l *Lexer) readString() string {
	var strBuilder strings.Builder
	for {
		l.readChar()
		if l.ch == '"' || l.ch == 0 {
			break
		}
		if l.ch == '\\' {
			l.readChar()
			switch l.ch {
			case 'n':
				strBuilder.WriteByte('\n')
			case 't':
				strBuilder.WriteByte('\t')
			case 'r':
				strBuilder.WriteByte('\r')
			case 'b':
				strBuilder.WriteByte('\b')
			case 'f':
				strBuilder.WriteByte('\f')
			case '\\':
				strBuilder.WriteByte('\\')
			case '"':
				strBuilder.WriteByte('"')
			default:
				strBuilder.WriteByte('\\')
				strBuilder.WriteByte(l.ch)
			}
		} else {
			strBuilder.WriteByte(l.ch)
		}
	}
	return strBuilder.String()
}

func (l *Lexer) readNumber() string {
	pos := l.position
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[pos:l.position]
}

func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func lookupIdent(ident string) token.TokenType {
	switch ident {
	case "true":
		return token.TRUE
	case "false":
		return token.FALSE
	case "null":
		return token.NULL
	}
	return token.ILLEGAL
}
