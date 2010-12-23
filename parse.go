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
