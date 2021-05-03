// Dup2 prints the count and text of lines that appear more than once in the input.
// It reads from stdin or from a list of named files.
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	counts := make(map[string]map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts, "Stdin")
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts, arg)
			f.Close()
		}
	}
	for line, nameCount := range counts {
		n := 0
		names, sep := "", ""
		for name, count := range nameCount {
			names += sep + name + "[" + strconv.Itoa(count) + "]"
			sep = " "
			n += count
		}
		if n > 1 {
			fmt.Printf("%d\t%s\t(%s)\n", n, line, names)
		}
		fmt.Println(counts)
	}
}

func countLines(f *os.File, counts map[string]map[string]int, name string) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		line := input.Text()
		if counts[line] == nil {
			counts[line] = make(map[string]int)
		}
		counts[line][name]++
	}
	// NOTE: ignoring potential errors from Input.Err()
}
