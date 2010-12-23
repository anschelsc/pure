package main

type Func interface {
	apply(Func) Func
	String() string
}

type raw string

func (r raw) String() string { return string(r) }

func (r raw) apply(arg Func) Func {
	return "`" + r + raw(arg.String())
}

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
	return "`" + c.first.String()
}

func (c s1) apply(second Func) Func { return s2{c.first, second} }

type s2 struct {
	first, second Func
}

func (c s2) String() string {
	return "``" + c.first.String() + c.second.String()
}

func (c s2) apply(third Func) Func {
	return c.first.apply(c.second).apply(c.first.apply(third))
}
