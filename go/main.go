package main

import (
	"fmt"
	"os"
	"io/ioutil"
	"flag"
	"log"
)

var (
	elim = flag.String("e", "", "Eliminate this (one-character) function.")
)

func main() {
	flag.Parse()
	var input []byte
	if flag.NArg() == 0 {
		var err os.Error
		input, err = ioutil.ReadAll(os.Stdin)
		if err != nil {
			log.Fatalln(err)
		}
	} else {
		file, err := os.Open(flag.Arg(0), os.O_RDONLY, 0)
		if err != nil {
			log.Fatalln(err)
		}
		defer file.Close()
		input, err = ioutil.ReadAll(file)
		if err != nil {
			log.Fatalln(err)
		}
	}
	input = strip(input)
	if !valid(input) {
		log.Fatalln("Syntax error.")
	}
	if len(*elim) == 0 {
		fmt.Println(parse(input))
	} else {
		input = []byte(parse(input).String())  //First we simplify.
		for i := len(*elim) - 1; i >= 0; i-- { //Iterate backwards.
			input = []byte(eliminate(dumbParse(input), char((*elim)[i])).String())
			input = []byte(parse(input).String()) //Simplify after every step.
		}
		fmt.Println(string(input))
	}
}
