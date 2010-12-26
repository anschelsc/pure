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

func eliminate(f simple, c char) Func {
	if fCh, ok := f.(char); ok && fCh == c {
		return I
	}
	if !f.contains(c) {
		return K.apply(f)
	}
	p := f.(pair) //safe since only pair and char implement simple, and char is always caught above
	if second, ok := p[1].(char); ok && second == c && !p[0].(simple).contains(c) {
		return p[0]
	}
	return S.apply(eliminate(p[0].(simple), c)).apply(eliminate(p[1].(simple), c))
}
