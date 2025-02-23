package parser

import (
	"fmt"
	"strings"
	"testing"
)

func TestParser(t *testing.T) {
	expression := `1234`
	parser := NewParser(strings.NewReader(expression))
	feelExpression, err := parser.ParseFeelExpression()
	if err != nil {
		t.Error(err)
	}
	fmt.Println(feelExpression)
}
