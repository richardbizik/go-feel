package main

import (
	"fmt"
	"testing"
)

func TestFeelParser(t *testing.T) {
	expression := "1+5"
	fmt.Println("starting")
	f := &Expression{
		Buffer: expression,
	}
	f.Init()
	if err := f.Parse(); err != nil {
		fmt.Println(err.Error())
	}
	fmt.Print(f)
}
