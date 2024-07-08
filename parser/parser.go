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

func (p *Parser) parseJSON() ast.Node {
	switch p.curToken.Type {
	case token.LBRACE:
		return p.parseObjectNode()
	}
	return nil
}

func (p *Parser) parseObjectNode() *ast.ObjectNode {
	var obj = &ast.ObjectNode{}
	if p.curToken.Type != token.LBRACE {
		panic("not valid JSON")
	}
  for p.curToken.Type != token.RBRACE{
    fmt.Println("looping inside an obj")
  }

	return nil
}
