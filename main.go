// dupgo prints the text of each line that appears more than
// once in the standard input, preceded by its count.
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]map[string]int)
	files := os.Args[1:]

	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)

			if err != nil {
				fmt.Fprintf(os.Stderr, "dupgo: %v\n", err)

				continue
			}

			defer f.Close()
			countLines(f, counts)
		}
	}

	for fName, m := range counts {
		hasDuplicates := false

		for line, n := range m {
			if n > 1 {
				if !hasDuplicates {
					fmt.Printf("\n%s: \n\n", fName)
				}
				hasDuplicates = true
				fmt.Printf("%d\t%s\n", n, line)
			}
		}
	}
}

func countLines(f *os.File, counts map[string]map[string]int) {
	input := bufio.NewScanner(f)

	for input.Scan() {
		if m, ok := counts[f.Name()]; ok {
			m[input.Text()]++
		} else {
			counts[f.Name()] = map[string]int{input.Text(): 1}
		}
	}

	// NOTE: ignoring potential errors from input.Err()
}
