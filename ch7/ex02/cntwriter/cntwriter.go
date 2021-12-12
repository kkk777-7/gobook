package cntwriter

import "io"

type Wrapper struct {
	w io.Writer
	c int64
}

func CountWriter(w io.Writer) (io.Writer, *int64) {
	var wc Wrapper
	wc.w = w
	return &wc, &(wc.c)
}

func (wc *Wrapper) Write(b []byte) (int, error) {
	n, err := wc.w.Write(b)
	if err != nil {
		return 0, err
	}
	wc.c += int64(n)
	return n, nil
}
