package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args
	fmt.Printf("The args are: %v\nArgument #1: %v\nAll arguments: %v\nLength of Args: %v\n", args, args[1], args[1:len(args)], len(args))
}
