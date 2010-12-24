package main

func valid(raw []byte) bool {
	count := 1
	for i := 0; i != len(raw); i++ {
		if count < 1 {
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
func split(raw []byte) int {
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
