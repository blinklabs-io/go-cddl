package main

import (
	"fmt"

	"github.com/blinklabs-io/go-cddl"
)

func main() {
	cddlStr := `
transaction_index = uint .size 2

next_major_protocol_version = 7
`
	parser := cddl.NewParser()
	ast, err := parser.ParseString("test", cddlStr)
	if err != nil {
		fmt.Printf("ERROR: %s\n", err)
		return
	}
	fmt.Printf("ast = %#v\n", ast)
}
