package main

import (
//	"os"
//	"io/ioutil"
	"fmt"
)

func main() {
//	raw, _ := ioutil.ReadAll(os.Stdin)
	raw := []byte("`f``nfx")
	p, e := Parse(raw)
	if e != nil {
		fmt.Println(e)
	}
	for _,k:=range []Char{'x','f','n'} {
		p=eliminate(p, k)
	}
	fmt.Println(p)
}
