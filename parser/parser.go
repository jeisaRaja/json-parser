package parser

import (
	"fmt"
	"jeisaraja/json_parser/ast"
	"jeisaraja/json_parser/lexer"
	"jeisaraja/json_parser/token"
	"strconv"
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

func (p *Parser) parseJSON() (ast.Node, error) {
	var result ast.Node
	switch p.curToken.Type {
	case token.LBRACE:
		result = p.parseObjectNode()
	case token.LBRACKET:
		p.parseArrayNode()
	default:
		p.error(fmt.Sprintf("Unexpected token %s", p.curToken.Literal))
	}
	if len(p.errors) > 0 {
		return nil, fmt.Errorf("parsing errors: %v", p.errors)
	}
	return result, nil
}

func (p *Parser) parseObjectNode() *ast.ObjectNode {
	var obj = &ast.ObjectNode{Pairs: []*ast.PairNode{}}
	if p.curToken.Type != token.LBRACE {
		p.error("Not correct")
		return obj
	}

	p.nextToken()
	fmt.Println("curToken is ", p.curToken.Type)
	for !p.peekTokenIs(token.RBRACE) {
		if !p.curTokenIs(token.STRING) {
			p.error("Expected string key")
			return obj
		}
		key := &ast.StringNode{Token: p.curToken, Value: p.curToken.Literal}
		fmt.Println("key is ", key.String())

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
		pair := &ast.PairNode{Key: key, Value: value}
		obj.Pairs = append(obj.Pairs, pair)
	}

	if p.curTokenIs(token.COMMA) {
		p.nextToken()
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

func (p *Parser) parseArrayNode() *ast.ArrayNode {
	node := &ast.ArrayNode{}
	if !p.curTokenIs(token.LBRACKET) {
		p.error(fmt.Sprintf("token is not correct, expected %s, got %s", p.curToken.Literal, token.LBRACKET))
	}

	for !p.peekTokenIs(token.RBRACKET) {
		p.nextToken()
		fmt.Println("hey in looping the array node")
		value := p.parseValue()
		if value == nil {
			return node
		}
		node.Value = append(node.Value, value)
		p.nextToken()
		if p.curTokenIs(token.COMMA) {
		} else if p.curTokenIs(token.RBRACKET) {
			break
		} else {
			fmt.Println("curToken is ", p.curToken)
			p.error("Expected ',' or ']' after value")
			return node
		}
	}

	if !p.curTokenIs(token.RBRACKET) {
		p.error("Expected ']' at end of array")
	}
	return node
}

func (p *Parser) parseValue() ast.Node {
	switch p.curToken.Type {
	case token.STRING:
		node := &ast.StringNode{Token: p.curToken, Value: p.curToken.Literal}
		return node
	case token.NUMBER:
		number, err := strconv.ParseFloat(p.curToken.Literal, 32)
		if err != nil {
			p.error(fmt.Sprintf("Failed to parse number for %s", p.curToken.Literal))
		}
		n := float32(number)
		return &ast.NumberNode{Token: p.curToken, Value: n}
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
