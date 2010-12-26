package main

func strip(raw []byte) []byte {
	ret := make([]byte, 0, len(raw))
	for i := 0; i != len(raw); i++ {
		switch raw[i] {
		case '\t', '\n', '\v', '\f', '\r', ' ', 0x85, 0xA0: //copied from unicode.IsSpace()
			continue
		default:
			ret = append(ret, raw[i])
		}
	}
	return ret
}

func valid(raw []byte) bool {
	count := 1
	for i := 0; i != len(raw); i++ {
		if count == 0 {
			return false
		}
		if raw[i] == '`' {
			count++
		} else {
			count--
		}
	}
	return count == 0
}

//Pass this a two-function slice; it returns the index of the border.
func border(raw []byte) int {
	count := 1
	for i := 0; i != len(raw); i++ {
		if count == 0 {
			return i
		}
		if raw[i] == '`' {
			count++
		} else {
			count--
		}
	}
	return -1
}

func split(raw []byte) ([]byte, []byte) {
	i := border(raw)
	return raw[:i], raw[i:]
}

//parse(raw) assumes that raw is valid().
func parse(raw []byte) Func {
	switch raw[0] {
	case '`':
		first, second := split(raw[1:])
		return parse(first).apply(parse(second))
	case 's':
		return S
	case 'k':
		return K
	case 'i':
		return I
	}
	return char(raw[0])
}
