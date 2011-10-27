package main

import (
	"os"
	"bufio"
	"fmt"
)

func main() {
	p, e := Parse(bufio.NewReader(os.Stdin))
	if e != nil {
		fmt.Fprintln(os.Stderr, e)
		return
	}
	fmt.Println(p.Eval().Defuse())
}
