package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {

	start := time.Now()

	for i := 0; i < 15000; i++ {
		/*
			s, sep := "", ""
			for _, arg := range os.Args[0:] {
				s += sep + arg
				sep = " "
			}
			fmt.Println(s)
		*/

		fmt.Println(strings.Join(os.Args[0:], " "))

	}

	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}
