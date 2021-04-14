// Dup2 prints the count and text of lines that appear more than once in the input.
// It reads from stdin or from a list of named files.
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	filenames := make(map[string]string)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts, filenames, "")
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts, filenames, arg)
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			var filenameList string
			for filename, _ := range filenames {
				filenameList += filename[line]
			}
			fmt.Printf("%d\t%s\t%s\n", n, line, filenameList)
		}
	}
}

func countLines(f *os.File, counts map[string]int, filenames map[string]string, arg string) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
		filenames[input.Text()] = arg
	}
	// NOTE: ignoring potential errors from Input.Err()
}
