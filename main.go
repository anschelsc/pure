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
			log.Exitln(err)
		}
	} else {
		file, err := os.Open(flag.Arg(0), os.O_RDONLY, 0)
		if err != nil {
			log.Exitln(err)
		}
		defer file.Close()
		input, err = ioutil.ReadAll(file)
		if err != nil {
			log.Exitln(err)
		}
	}
	input = strip(input)
	if !valid(input) {
		log.Exitln("Syntax error.")
	}
	switch len(*elim) {
	case 0:
		fmt.Println(parse(input))
	case 1:
		fmt.Println(eliminate(dumbParse(input), char((*elim)[0])))
	default:
		log.Exitln("Argument to -e should be one-character.")
	}
}
