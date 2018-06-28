package main

import (
	"fmt"
	"os"
)

func main() {

	// get the command line args
	args := os.Args[1:]
	fmt.Println(args)
}
