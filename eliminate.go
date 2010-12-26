package main

//simple is a simple function, with no combinators.
type simple interface {
	Func
	contains(char) bool
}

func (c char) contains(other char) bool { return c == other }

func (p pair) contains(c char) bool {
	var first, second bool
	if s, ok := p[0].(simple); ok {
		first = s.contains(c)
	}
	if s, ok := p[1].(simple); ok {
		second = s.contains(c)
	}
	return first || second
}

func dumbParse(raw []byte) simple {
	if raw[0] == '`' {
		first, second := split(raw[1:])
		return pair{dumbParse(first), dumbParse(second)}
	}
	return char(raw[0])
}
