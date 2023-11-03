package main

import (
	"fmt"
	"os"
)

// Args is a wrapper around os.Args
// that provides some convenience methods
// for working with command line arguments
// without having to worry about bounds checking
type Args []string

// Num returns the string at the given index,
// or an empty string if the index is out of bounds
func (arg Args) Num(i int) string {
	if i < 0 || i >= len(arg) {
		return ""
	}
	return arg[i]
}

// ParseArgs parses the command line arguments to Args
// skip the first argument, which is the program name
func ParseArgs() Args {
	if len(os.Args) > 1 {
		return os.Args[1:]
	}

	return Args{}
}

func main() {
	// take an os.Args and wrap it in our Args type
	args := ParseArgs()

	fmt.Println("args:", os.Args)

	// use the new methods
	fmt.Println("arg 1:", args.Num(0))
}
