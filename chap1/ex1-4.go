package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	where := make(map[string]string)
	files := os.Args[1:]
	if len(files) == 0 {
		countLinesEx(os.Stdin, counts, where)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLinesEx(f, counts, where)
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
			fmt.Printf(where[line] + "\n")
		}
	}
}

func countLinesEx(f *os.File, counts map[string]int, where map[string]string) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
		where[input.Text()] += " " + f.Name()
	}
}
