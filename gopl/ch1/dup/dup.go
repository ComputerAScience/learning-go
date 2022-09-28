package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counters := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		counters[input.Text()]++
	}

	for line, count := range counters {
		if count > 1 {
			fmt.Printf("%s\t%d\n", line, count)
		}
	}
}
