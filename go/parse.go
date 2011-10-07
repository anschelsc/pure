package main

import (
	"os"
	"bytes"
)

var (
	ErrSyntax = os.NewError("Syntax error.")
)

var whiteSpace = map[byte]bool{
	'\t': true,
	'\n': true,
	'\v': true,
	'\f': true,
	'\r': true,
	' ':  true,
	0x85: true,
	0xA0: true,
}

func Parse(s []byte) (p Piece, e os.Error) {
	defer func() {
		x := recover()
		if err, ok := x.(os.Error); ok {
			p = nil
			e = err
		} else if x != nil {
			panic(x)
		}
	}()
	p, rest := parse(s)
	if len(rest) != 0 && len(bytes.TrimSpace(rest)) != 0 {
		e = ErrSyntax
	}
	return
}

func parse(s []byte) (Piece, []byte) {
	if len(s) == 0 {
		panic(ErrSyntax)
	}
	if whiteSpace[s[0]] {
		return parse(s[1:])
	}
	if s[0] == '`' {
		left, s := parse(s[1:])
		right, s := parse(s)
		return Pair{left, right}, s
	}
	return Char(s[0]), s[1:]
}
