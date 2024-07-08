package parser

import (
	"fmt"
	"jeisaraja/json_parser/lexer"
	"testing"
)

func TestParser(t *testing.T) {
	input := "{}"
	l := lexer.New(input)
	p := New(l)

	result := p.parseJSON()

	if result == -1 {
    fmt.Print(p.errors)
		t.Fatalf("parseJSON() returned -1 for input: %s", input)
	} else {
		return
	}
}
