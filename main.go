package main

import (
	"fmt"
)

func main() {
	fmt.Println(char('a').apply(char('b')))
	fmt.Println(char('a').apply(I))
	fmt.Println(I.apply(char('a')))
	K1 := K.apply(char('a'))
	fmt.Println(K1)
	fmt.Println(K1.apply(char('b')))
	S1 := S.apply(char('a'))
	S2 := S1.apply(char('b'))
	fmt.Println(S1)
	fmt.Println(S2)
	fmt.Println(S2.apply(char('c')))
}
