package main

import (
	"fmt"
	"os"
)

func main() {
	s, sep := "", " "
	for idx, arg := range os.Args[1:] {
		s += fmt.Sprint(idx) + sep + arg + "\n"
	}
	fmt.Print(s)
}
