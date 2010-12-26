package main

func dumbParse(raw []byte) Func {
	if raw[0] == '`' {
		first, second := split(raw[1:])
		return dumbParse(first).apply(dumbParse(second))
	}
	return char(raw[0])
}

//f should be the result of dumbParse, i.e. it doesn't understand combinators.
func eliminate(f Func, c char) Func {
	switch v := f.(type) {
	case char:
		if v == c {
			return I
		}
		return K.apply(v)
	case pair:
		if v.contains(c) {
			//check if v is `<something>c
			if second, ok := v[1].(char); ok {
				if first, ok := v[0].(pair); ok {
					if second == c && !first.contains(c) {
						return first
					}
				} else {
					if second == c && v[0].(char) != c {
						return first
					}
				}
			}
			return S.apply(eliminate(v[0], c)).apply(eliminate(v[1], c))
		}
		return K.apply(v)
	}
	return nil
}

func (p pair) contains(c char) bool {
	if test, ok := p[0].(char); ok {
		if test == c {
			return true
		}
	} else {
		if p[0].(pair).contains(c) {
			return true
		}
	}
	//p[0] does not contain c
	if test, ok := p[1].(char); ok {
		if test == c {
			return true
		}
		return false
	}
	return p[1].(pair).contains(c)
}
