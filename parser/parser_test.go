package parser

import (
	"fmt"
	"jeisaraja/json_parser/lexer"
	"testing"
)

func TestParser(t *testing.T) {
	testCases := []struct {
		input          string
		expectedOutput interface{}
		shouldFail     bool
	}{
		// Valid JSON objects
		{"{\"key\": \"value\"}", map[string]interface{}{"key": "value"}, false},
		{"{\"num\": 123}", map[string]interface{}{"num": 123}, false},
		{"{\"bool\": true}", map[string]interface{}{"bool": true}, false},
		{"{\"null\": null}", map[string]interface{}{"null": nil}, false},
		{"{\"array\": [1, 2, 3]}", map[string]interface{}{"array": []interface{}{1, 2, 3}}, false},
		{"{\"nested\": {\"key\": \"value\"}}", map[string]interface{}{"nested": map[string]interface{}{"key": "value"}}, false},

		// Invalid JSON objects
		{"{\"unclosedKey\": \"value\"", nil, true},
		{"{\"key\": \"unclosedValue}", nil, true},
		{"{\"missingColon\" \"value\"}", nil, true},
		{"{\"extraComma\": \"value\",}", nil, true},
		{"{\"key\": \"value\", \"key2\":}", nil, true},

		// Valid JSON arrays
		{"[\"value1\", \"value2\"]", []interface{}{"value1", "value2"}, false},
		{"[123, 456]", []interface{}{123, 456}, false},
		{"[true, false]", []interface{}{true, false}, false},
		{"[null, \"string\", 123]", []interface{}{nil, "string", 123}, false},
		{"[{\"key\": \"value\"}, {\"key2\": \"value2\"}]", []interface{}{map[string]interface{}{"key": "value"}, map[string]interface{}{"key2": "value2"}}, false},

		// Invalid JSON arrays
		{"[\"unclosedValue]", nil, true},
		{"[\"value1\", \"value2\",]", nil, true},
		{"[true,]", nil, true},
		{"[, \"value\"]", nil, true},

		// Edge cases
		{"{}", map[string]interface{}{}, false},
		{"[]", []interface{}{}, false},
		{"{\"emptyString\": \"\"}", map[string]interface{}{"emptyString": ""}, false},
		{"[null]", []interface{}{nil}, false},
	}

	for _, tc := range testCases {
		l := lexer.New(tc.input)
		p := New(l)

		result := p.parseJSON()

		if tc.shouldFail {
			if result != -1 {
				t.Fatalf("Expected failure for input: %s, but got success", tc.input)
			} else {
				fmt.Printf("Correctly failed for input: %s with errors: %v\n", tc.input, p.errors)
			}
		} else {
			if result == -1 {
				fmt.Print(p.errors)
				t.Fatalf("parseJSON() returned -1 for input: %s", tc.input)
			} else {
				fmt.Printf("Successfully parsed input: %s with result: %v\n", tc.input, result)
			}
		}
	}
}

