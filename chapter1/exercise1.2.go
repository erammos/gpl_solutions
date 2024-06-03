package main

import (
	"fmt"
	"os"
)

func main() {

	var s string
	for i := 1; i < len(os.Args); i++ {
		s += string(i+'0') + " " + os.Args[i] + "\n"
	}
	fmt.Println(s)

}
