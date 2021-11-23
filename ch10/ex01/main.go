package main

import (
	"flag"
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"
	"os"
)

var f string

func init() {
	flag.StringVar(&f, "F", "png", "Encode format")
}
func main() {
	flag.Parse()
	if err := changeFormat(os.Stdin, os.Stdout, f); err != nil {
		fmt.Fprintf(os.Stderr, "jpeg: %v\n", err)
		os.Exit(1)
	}
}

func changeFormat(in io.Reader, out io.Writer, format string) error {
	img, kind, err := image.Decode(in)
	if err != nil {
		return err
	}
	fmt.Fprintln(os.Stderr, "Input format =", kind)
	switch f {
	case "png":
		return png.Encode(out, img)
	case "jpeg":
		return jpeg.Encode(out, img, &jpeg.Options{Quality: 95})
	case "gif":
		return gif.Encode(out, img, nil)
	}
	return nil
}
