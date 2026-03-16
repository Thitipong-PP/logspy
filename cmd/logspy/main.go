package main

import (
	"fmt"
	"os"

	"github.com/Thitipong-PP/logspy/internal/parser"
)

func main() {
	// Use parser for cutting os arguments to statement
	statement, err := parser.Parse(os.Args)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	fmt.Println("Statement: ", statement)
}
