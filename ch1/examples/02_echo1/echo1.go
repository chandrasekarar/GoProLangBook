// Echo1 prints its command-line arguments.
package main

import (
	"fmt"
	"os"
)

func main() {
	var s, sep string

	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}

	fmt.Println(s)

	//below code prints path to file
	// fmt.Println(os.Args[0])
}

//to run
//go run echo1.go hello world
// prints hello world
