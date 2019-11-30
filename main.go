package main

import (
	"fmt"
	"monkeylang/repl"
	"os"
)

func main() {
	fmt.Printf("Welcome to monkeylang!\nThis prompt now accepts your commands:\n")

	repl.Start(os.Stdin, os.Stdout)
}
