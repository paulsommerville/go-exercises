// Echo2 prints its command-line arguments.
package main

import (
	"fmt"
	"os"
)

func main() {
	var line, s string
	for i, arg := range os.Args[1:] {
		line = fmt.Sprint(i, " ", arg, " GIT\n")
		s += line
	}
	fmt.Println(s)
}
