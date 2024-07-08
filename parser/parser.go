package parser

import (
	"fmt"
	"jeisaraja/json_parser/ast"
	"jeisaraja/json_parser/lexer"
	"jeisaraja/json_parser/token"
)

type Parser struct {
	lexer *lexer.Lexer

	curToken  token.Token
	peekToken token.Token

	errors []string
}

func New(l *lexer.Lexer) *Parser {
	parser := &Parser{lexer: l}
	parser.nextToken()
	parser.nextToken()
	return parser
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.lexer.NextToken()
}

func (p *Parser) parseJSON() int {
	switch p.curToken.Type {
	case token.LBRACE:
		p.parseObjectNode()
	}
	if len(p.errors) > 0 {
		return -1
	}
	return 0
}

func (p *Parser) parseObjectNode() *ast.ObjectNode {
	var obj = &ast.ObjectNode{}
	if p.curToken.Type != token.LBRACE {
		p.error("Not correct")
	}
	for p.curToken.Type != token.RBRACE {
		p.nextToken()
		fmt.Println("looping inside an obj")
	}

	if !p.peekTokenIs(token.EOF) {
		p.error("Not Valid, something after end of closing brace")
		return obj
	}
	return obj
}

func (p *Parser) peekTokenIs(t token.TokenType) bool {
	return t == p.peekToken.Type
}

func (p *Parser) curTokenIs(t token.TokenType) bool {
	return t == p.curToken.Type
}

func (p *Parser) error(s string) {
	p.errors = append(p.errors, s)
}
