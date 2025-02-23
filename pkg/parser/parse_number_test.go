package parser

import (
	"fmt"
	"testing"
)

func TestParseNumber(t *testing.T) {
	intNumber := parseNumber("1234")
	if intNumber != int64(1234) {
		t.Error("failed to parse number")
	}

	intNumber = parseNumber("1")
	if intNumber != int64(1) {
		t.Error("failed to parse number")
	}

	floatNumber := parseNumber("1234.567")
	if floatNumber != 1234.567 {
		t.Error("failed to parse number")
	}

	intNumber = parseNumber("1234.")
	if intNumber != int64(1234) {
		t.Error("failed to parse number")
	}

	expNumber := parseNumber("12.2e2")
	expected := float64(12.2) * float64(12.2)
	if expNumber != expected {
		t.Error("failed to parse number")
	}

	expNumber = parseNumber("10e2")
	expected = float64(10) * float64(10)
	if expNumber != expected {
		t.Error("failed to parse number")
	}
}

func TestK(t *testing.T) {
	fmt.Println(10 ^ 2)

}
