package main

import (
	"fmt"
)

var test = []string{"a", "`ab", "`ka", "``kab", "`ix", "```sxyz"}

func main() {
	for _, s := range test {
		fmt.Printf("%s\t%s\n", s, parse([]byte(s)))
	}
}
