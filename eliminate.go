package main

//simple is a simple function, with no combinators.
type simple interface {
	Func
	contains(char) bool
}

func (c char) contains(other char) bool { return c == other }

func (p pair) contains(c char) bool {
	return p[0].(simple).contains(c) || p[1].(simple).contains(c)
}

func dumbParse(raw []byte) simple {
	if raw[0] == '`' {
		first, second := split(raw[1:])
		return pair{dumbParse(first), dumbParse(second)}
	}
	return char(raw[0])
}
