package main

import (
	"fmt"
)

type Piece interface {
	Eval() Func
	String() string
}

type Func interface {
	Apply(Piece) Func
	Defuse() Piece
}

type Char byte

func (c Char) String() string { return string(c) }

func (c Char) Eval() Func {
	switch c {
	case 's':
		return S{}
	case 'k':
		return K{}
	case 'i':
		return I{}
	}
	return Block{c}
}

type Pair [2]Piece

func (p Pair) String() string { return fmt.Sprintf("`%s%s", p[0], p[1]) }

func (p Pair) Eval() Func { return p[0].Eval().Apply(p[1]) }

type S struct{}

func (s S) Apply(p Piece) Func { return S1{p} }

func (s S) Defuse() Piece { return Char('s') }

type S1 struct {
	x Piece
}

func (s S1) Apply(y Piece) Func { return S2{s.x, y} }

func (s S1) Defuse() Piece { return Pair{Char('s'), s.x} }

type S2 struct {
	x, y Piece
}

func (s S2) Apply(z Piece) Func { return s.x.Eval().Apply(z).Apply(Pair{s.y, z}) }

func (s S2) Defuse() Piece { return Pair{Pair{Char('s'), s.x}, s.y} }

type K struct{}

func (k K) Apply(p Piece) Func { return K1{p} }

func (k K) Defuse() Piece { return Char('k') }

type K1 struct {
	p Piece
}

func (k K1) Apply(_ Piece) Func { return k.p.Eval() }

func (k K1) Defuse() Piece { return Pair{Char('k'), k.p} }

type I struct{}

func (i I) Apply(p Piece) Func { return p.Eval() }

func (i I) Defuse() Piece { return Char('i') }

type Block struct {
	p Piece
}

func (b Block) Apply(p Piece) Func { return Block{Pair{b.p, p}} }

func (b Block) Defuse() Piece { return b.p }
