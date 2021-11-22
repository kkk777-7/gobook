package main

import (
	"bufio"
	"fmt"
	"gobook/ch2/ex02/conv"
	"os"
	"strconv"
)

func main() {
	args := getArgs()
	for _, arg := range args {
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cf: %v\n", err)
			os.Exit(1)
		}
		conv.Temp(t)
		conv.Length(t)
		conv.Weight(t)
		fmt.Println("-----")
	}
}

func getArgs() []string {
	var args []string
	if len(os.Args) == 1 {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			arg := scanner.Text()
			if arg == "end" {
				fmt.Println("-----")
				break
			}
			args = append(args, arg)
		}
	} else {
		args = os.Args[1:]
	}
	return args
}
