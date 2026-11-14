func canConstruct(ransomNote string, magazine string) bool {
	var count [26]int
	for _, c := range magazine {
		count[c-'a']++
	}
	for _, c := range ransomNote {
		count[c-'a']--
		if count[c-'a'] < 0 {
			return false
		}
	}

	return true
}
