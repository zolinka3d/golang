func updateKey(key []byte, c1, c2, c3 []byte) []byte {
	min := MinIntSlice([]int{len(key), len(c1), len(c2), len(c3)})
	r1 := bytes.Runes(c1)
	r2 := bytes.Runes(c2)
	r3 := bytes.Runes(c3)
	space := byte(' ')
	for i := 0; i < min; i++ {
		if key[i] != '#' || r1[i] == r2[i] || r1[i] == r3[i] || r2[i] == r3[i] {
			continue
		}
		r12 := r1[i] ^ r2[i]
		r13 := r1[i] ^ r3[i]
		r23 := r2[i] ^ r3[i]

		if unicode.IsLetter(r12) && unicode.IsLetter(r13) {
			key[i] = c1[i] ^ space
		} else if unicode.IsLetter(r12) && unicode.IsLetter(r23) {
			key[i] = c2[i] ^ space
		} else if unicode.IsLetter(r13) && unicode.IsLetter(r23) {
			key[i] = c3[i] ^ space
		}
	}

	return key
}