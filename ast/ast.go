package ast

import (
	"fmt"
	"jeisaraja/json_parser/token"
)

type Node interface {
	TokenLiteral() string
	String() string
}

type ObjectNode struct {
	Pairs []*PairNode
}

type ArrayNode struct {
	Value []Node
}

type BooleanNode struct {
	Token token.Token
	Value bool
}

type NullNode struct {
	Token token.Token
}

func (on *ObjectNode) String() string {
	return ""
}

func (on *ObjectNode) TokenLiteral() string {
	return ""
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
	return fmt.Sprintf("StringNode{Token: %s, Value: %s}", sn.Token.Type, sn.Value)
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

func (bn *BooleanNode) String() string {
	return bn.Token.Literal
}

func (bn *BooleanNode) TokenLiteral() string {
	return bn.Token.Literal
}

func (nn *NullNode) String() string {
	return nn.Token.Literal
}

func (nn *NullNode) TokenLiteral() string {
	return nn.Token.Literal
}

func (an *ArrayNode) String() string {
	return fmt.Sprintf("%s", an.Value)
}

func (an *ArrayNode) TokenLiteral() string {
	return fmt.Sprintf("%s", an.Value)
}
