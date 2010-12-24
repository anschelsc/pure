package main

import (
	"fmt"
)

var test = []string{"aa", "`aaa", "a`aa"}

func main() {
	for _, s := range test {
		fmt.Printf("%s\t%d\n", s, split([]byte(s)))
	}
}
