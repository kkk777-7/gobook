package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string)
	file, err := os.OpenFile("test.txt", os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	for _, url := range os.Args[1:] {
		go fetch(url, ch)
	}
	for range os.Args[1:] {
		res := <-ch
		res += "\n"
		resbyte := bytes.NewBufferString(res)
		_, err := file.Write(resbyte.Bytes())
		if err != nil {
			log.Fatal(err)
		}
	}
	fin := fmt.Sprintf("%.2fs elapsed\n-----\n", time.Since(start).Seconds())
	finbyte := bytes.NewBufferString(fin)
	_, err = file.Write(finbyte.Bytes())
	if err != nil {
		log.Fatal(err)
	}
}

func fetch(url string, ch chan<- string) {
	if !strings.HasPrefix(url, "http://") {
		url = "http://" + url
	}
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}

	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs  %7d  %s", secs, nbytes, url)
}
