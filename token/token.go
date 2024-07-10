package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	LBRACE   = "{"
	RBRACE   = "}"
	LBRACKET = "["
	RBRACKET = "]"
	COLON    = ":"
	COMMA    = ","

	NUMBER = "NUMBER"
	STRING = "STRING"

	TRUE  = "TRUE"
	FALSE = "FALSE"
	NULL  = "NULL"
)
