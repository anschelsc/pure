package main

func contains(p Piece, x Char) bool {
	switch pp := p.(type) {
	case Char:
		return x == pp
	case Pair:
		return contains(pp[0], x) || contains(pp[1], x)
	}
	panic("unreachable")
}

func eliminate(p Piece, x Char) Piece {
	if !contains(p, x) {
		return Pair{Char('k'), p}
	}
	switch pp := p.(type) {
	case Char: // pp==x
		return Char('i')
	case Pair:
		if c, ok := pp[1].(Char); ok {
			if c == x && !contains(pp[0], x) {
				return pp[0]
			}
		}
		return Pair{Pair{Char('s'), eliminate(pp[0], x)}, eliminate(pp[1], x)}
	}
	panic("unreachable")
}
