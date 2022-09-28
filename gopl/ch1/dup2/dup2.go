package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	counters := make(map[string]int)
	if len(os.Args) == 1 {
		countLines(os.Stdin, counters)
	} else {
		for _, arg := range os.Args[1:] {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counters)
			f.Close()
		}
	}

	for line, count := range counters {
		if count > 1 {
			fmt.Printf("%s\t%d\n", line, count)
		}
	}
}

func countLines(f *os.File, c map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		c[input.Text()]++
	}
}
