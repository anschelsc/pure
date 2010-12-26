package main

//simple is a simple function, with no combinators.
type simple interface {
	Func
	contains(char) bool
}

//Make char implement simple.
func (c char) contains(other char) bool { return c == other }

type simplePair [2]simple

func (p simplePair) String() string {
	return "`" + p[0].String() + p[1].String()
}

func (p simplePair) apply(arg Func) Func {
	if s, ok := arg.(simple); ok { //Keep the simpleness if possible
		return simplePair{p, s}
	}
	return pair{p, arg}
}

func (p simplePair) contains(c char) bool {
	return p[0].contains(c) || p[1].contains(c)
}

func dumbParse(raw []byte) simple {
	if raw[0] == '`' {
		first, second := split(raw[1:])
		return simplePair{dumbParse(first), dumbParse(second)}
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
	p := f.(simplePair) //safe since only simplePair and char implement simple, and char is always caught above
	if second, ok := p[1].(char); ok && second == c && !p[0].contains(c) {
		return p[0]
	}
	return S.apply(eliminate(p[0], c)).apply(eliminate(p[1], c))
}
