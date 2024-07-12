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
	case token.LBRACKET:
		p.parseArrayObject()
	default:
		p.error(fmt.Sprintf("Unexpected token %s", p.curToken.Literal))
	}
	if len(p.errors) > 0 {
		return -1
	}
	return 0
}

func (p *Parser) parseObjectNode() *ast.ObjectNode {
	var obj = &ast.ObjectNode{Pairs: []*ast.PairNode{}}
	if p.curToken.Type != token.LBRACE {
		p.error("Not correct")
		return obj
	}

	for !p.peekTokenIs(token.RBRACE) {
		p.nextToken()
		if !p.curTokenIs(token.STRING) {
			p.error("Expected string key")
			return obj
		}
		key := &ast.StringNode{Token: p.curToken, Value: p.curToken.Literal}
		p.nextToken()

		if !p.curTokenIs(token.COLON) {
			p.error("Expected ':' after key")
			return obj
		}
		p.nextToken()
		value := p.parseValue()
		if value == nil {
			return obj
		}
	}

	if !p.peekTokenIs(token.EOF) {
		p.error("Not Valid, something after end of closing brace")
		return obj
	}
	return obj
}

func (p *Parser) peekTokenIs(t token.TokenType) bool {
	fmt.Println(p.peekToken)
	return t == p.peekToken.Type
}

func (p *Parser) curTokenIs(t token.TokenType) bool {
	return t == p.curToken.Type
}

func (p *Parser) error(s string) {
	p.errors = append(p.errors, s)
}

func (p *Parser) parseArrayObject() {

}

func (p *Parser) parseValue() ast.Node {
	switch p.curToken.Type {
	case token.STRING:
		return &ast.StringNode{Value: p.curToken.Literal}
	case token.NUMBER:
		return &ast.NumberNode{Value: p.curToken.Literal}
	case token.TRUE, token.FALSE:
		return &ast.BooleanNode{Value: p.curToken.Literal == "true"}
	case token.NULL:
		return &ast.NullNode{}
	case token.LBRACE:
		return p.parseObjectNode()
	case token.LBRACKET:
		return p.parseArrayNode()
	default:
		p.error(fmt.Sprintf("Unexpected token: %s", p.curToken.Literal))
		return nil
	}
}
