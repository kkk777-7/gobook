package reverse

const Size int = 10

func Do(s *[Size]int) {
	for i, j := 0, Size-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
