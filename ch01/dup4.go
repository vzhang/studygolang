// exercise 1.4 Modify dup2 to print the names of all files in which each duplicated line occurs.
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
		countLinesWithFilename(os.Stdin, "os.Stdin", counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup4: %v\n", err)
				continue
			}

			countLinesWithFilename(f, arg, counts)
			f.Close()
		}
	}

	for filename, innerMap := range counts {
		for line, n := range innerMap {
			if n > 1 {
				fmt.Printf("%s\t%s\t%d\t\n", filename, line, n)
			}
		}
	}
}

func countLinesWithFilename(f *os.File, filename string, counts map[string]map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		if counts[filename] == nil {
			counts[filename] = make(map[string]int)
		}

		counts[filename][input.Text()]++
	}
}
