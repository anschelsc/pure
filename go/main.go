package main

import (
	"os"
	"io/ioutil"
	"fmt"
)

func main() {
	raw, _ := ioutil.ReadAll(os.Stdin)
	p,e := Parse(raw)
	if e!=nil {
		fmt.Println(e)
		return
	}
	fmt.Println(p.Eval().Defuse())
}
