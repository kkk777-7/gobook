package rotate

func Do(s []int, n int) []int {
	rotate := s[:n]
	s = s[n:]
	s = append(s, rotate...)
	return s
}
