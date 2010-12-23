package main

import (
	"fmt"
)

func main() {
	fmt.Println(raw("a").apply(raw("b")))
	fmt.Println(raw("a").apply(I))
	fmt.Println(I.apply(raw("a")))
	K1 := K.apply(raw("a"))
	fmt.Println(K1)
	fmt.Println(K1.apply(raw("b")))
	S1 := S.apply(raw("a"))
	S2 := S1.apply(raw("b"))
	fmt.Println(S1)
	fmt.Println(S2)
	fmt.Println(S2.apply(raw("c")))
}
