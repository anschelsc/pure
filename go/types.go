package main

type Func interface {
	apply(Func) Func
	String() string
}

type char byte

func (c char) String() string { return string(c) }

func (c char) apply(arg Func) Func { return pair{c, arg} }

type pair [2]Func

func (p pair) String() string {
	return "`" + p[0].String() + p[1].String()
}

func (p pair) apply(arg Func) Func { return pair{p, arg} }

type combinator uint8

const (
	S combinator = iota
	K
	I
)

func (c combinator) String() string {
	switch c {
	case S:
		return "s"
	case K:
		return "k"
	case I:
		return "i"
	}
	return ""
}

func (c combinator) apply(arg Func) Func {
	switch c {
	case I:
		return arg
	case K:
		return k1{arg}
	case S:
		return s1{arg}
	}
	return nil
}

type k1 struct {
	arg Func
}

func (c k1) String() string {
	return "`k" + c.arg.String()
}

func (c k1) apply(_ Func) Func { return c.arg }

type s1 struct {
	first Func
}

func (c s1) String() string {
	return "`s" + c.first.String()
}

func (c s1) apply(second Func) Func { return s2{c.first, second} }

type s2 struct {
	first, second Func
}

func (c s2) String() string {
	return "``s" + c.first.String() + c.second.String()
}

func (c s2) apply(third Func) Func {
	return c.first.apply(third).apply(c.second.apply(third))
}
