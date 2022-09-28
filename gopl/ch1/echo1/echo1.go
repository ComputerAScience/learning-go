// Echo1 prints its command-line arguments.
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)

	s, sep = "", ""
	for _, i := range os.Args[1:] {
		s += sep + i
		sep = " "
	}
	fmt.Println(s)

	fmt.Println(strings.Join(os.Args[1:], " "))

	reader := bufio.NewReader(os.Stdin)
	bytes := make([]byte, 100)
	n, _ := reader.Read(bytes)
	fmt.Printf("%s", bytes[:n])
}
