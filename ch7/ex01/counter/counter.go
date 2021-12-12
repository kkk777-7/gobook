package counter

import (
	"bufio"
)

type WordCounter int
type LineCounter int

func (c *WordCounter) Write(p []byte) (int, error) {
	var cnt int

	for len(p) > 0 {
		n, _, err := bufio.ScanWords(p, true)
		if err != nil {
			return 0, err
		}
		p = p[n:]
		cnt++
	}
	*c += WordCounter(cnt)
	return cnt, nil
}

func (c *LineCounter) Write(p []byte) (int, error) {
	var cnt int

	for len(p) > 0 {
		n, _, err := bufio.ScanLines(p, true)
		if err != nil {
			return 0, err
		}
		p = p[n:]
		cnt++
	}
	*c += LineCounter(cnt)
	return cnt, nil
}
