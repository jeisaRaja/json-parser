package ast

import "jeisaraja/json_parser/token"

type Node interface {
	TokenLiteral() string
	String() string
}

type ObjectNode struct {
	Pairs []*PairNode
}

type PairNode struct {
	Key   *StringNode
	Value Node
}

type StringNode struct {
	Token token.Token
	Value string
}

func (sn *StringNode) TokenLiteral() string {
	return sn.Token.Literal
}

func (sn *StringNode) String() string {
	return sn.Token.Literal
}

type NumberNode struct {
	Token token.Token
	Value float32
}
func (nn *NumberNode) TokenLiteral() string {
	return nn.Token.Literal
}

func (nn *NumberNode) String() string {
	return nn.Token.Literal
}
