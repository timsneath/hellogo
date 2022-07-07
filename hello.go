package main

import (
	"fmt"

	"github.com/google/go-cmp/cmp"
	"github.com/timsneath/hellogo/morestrings"
)

func main() {
	// fmt.Println("Hello again")
	fmt.Println(morestrings.ReverseRunes("!oG ,olleH"))
	fmt.Println(cmp.Diff("Hello world", "Hello World"))
}
