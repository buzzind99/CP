func isIsomorphic(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	sToT := make(map[byte]byte)
	tToS := make(map[byte]byte)
	for i := range len(s) {
		sChar, tChar := s[i], t[i]
		if mapped, exists := sToT[sChar]; exists {
			if mapped != tChar {
				return false
			}
		} else {
			sToT[sChar] = tChar
		}

		if mapped, exists := tToS[tChar]; exists {
			if mapped != sChar {
				return false
			}
		} else {
			tToS[tChar] = sChar
		}
	}

	return true
}
