package main

import (
	"fmt"
	"os"

	"github.com/alecthomas/participle/v2/ebnf"
)

func main2() {
	f, err := os.Open("ebnf")
	if err != nil {
		panic(err)
	}
	p, err := ebnf.Parse(f)
	if err != nil {
		panic(err)
	}

	fmt.Println(p)

}
