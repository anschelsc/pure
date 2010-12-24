package main

import (
	"fmt"
	"os"
	"io/ioutil"
	"flag"
	"log"
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
	parsed := parse(input)
	fmt.Println(parsed)
}
