package main

import (
	"fmt"
)

var test = []string{"a", "`ab", "``aaa", "`", "a`", "a`a"}

func main() {
	for _, s := range test {
		fmt.Printf("%s\t%t\n", s, valid([]byte(s)))
	}
}
