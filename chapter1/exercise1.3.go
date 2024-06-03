package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func concat_method() {

	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}
func join_method() {
	fmt.Println(strings.Join(os.Args[1:], " "))
}

func main() {

	start := time.Now()
	concat_method()
	secs := time.Since(start)

	fmt.Println("Concat time:", secs)

	start = time.Now()
	join_method()
	secs = time.Since(start)

	fmt.Println("Join time:", secs)
}
