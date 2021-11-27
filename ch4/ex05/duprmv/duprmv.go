package duprmv

func Do(s []string) []string {
	var i int
	var chk bool
	for !chk {
		if s[i] == s[i+1] {
			copy(s[i:], s[i+1:])
			s = s[:len(s)-1]
			i--
		}
		i++
		if i == len(s)-1 {
			chk = true
		}
	}
	return s
}
