package main

import (
	"errors"
	"io"
)

var (
	ErrSyntax = errors.New("Syntax error.")
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

func Parse(r io.ByteReader) (Piece, error) {
	b, err := r.ReadByte()
	for err == nil && whiteSpace[b] {
		b, err = r.ReadByte()
	}
	if err != nil {
		if err == io.EOF {
			return nil, ErrSyntax
		}
		return nil, err
	}
	if b == '`' {
		left, err := Parse(r)
		if err != nil {
			return nil, err
		}
		right, err := Parse(r)
		if err != nil {
			return nil, err
		}
		return Pair{left, right}, nil
	}
	return Char(b), nil
}
