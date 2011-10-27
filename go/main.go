package main

import (
	"os"
	"bufio"
	"fmt"
	"flag"
)

var (
	elims = flag.String("e", "", "Eliminate these (one-character) functions.")
	input *os.File
)

func main() {
	flag.Parse()
	if flag.NArg() == 0 {
		input = os.Stdin
	} else {
		var err os.Error
		input, err = os.Open(flag.Arg(0))
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return
		}
	}
	p, e := Parse(bufio.NewReader(input))
	if e != nil {
		fmt.Fprintln(os.Stderr, e)
		return
	}
	for i := len(*elims) - 1; i >= 0; i-- { // NOP if *elims==""
		p = eliminate(p, Char((*elims)[i]))
	}
	fmt.Println(p.Eval().Defuse())
}
