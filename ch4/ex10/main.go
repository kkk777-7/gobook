package main

import (
	"fmt"
	"gobook/ch4/ex10/github"
	timeutil "gobook/ch4/ex10/timeUtil"
	"log"
	"os"
	"time"
)

func main() {
	var day []*github.Issue
	var month []*github.Issue
	var year []*github.Issue

	now := time.Now().UTC()
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d issues:\n", result.TotalCount)
	for _, item := range result.Items {
		if timeutil.LessMonth(item.CreatedAt, now) {
			day = append(day, item)
			continue
		}
		if timeutil.LessYear(item.CreatedAt, now) {
			month = append(month, item)
			continue
		}
		year = append(year, item)
	}
	fmt.Println("--- less than one month ---")
	for _, item := range day {
		fmt.Printf("#%-5d %9.9s %.55s\n",
			item.Number, item.User.Login, item.Title)
	}

	fmt.Println("--- less than one year ---")
	for _, item := range month {
		fmt.Printf("#%-5d %9.9s %.55s\n",
			item.Number, item.User.Login, item.Title)
	}

	fmt.Println("--- more than one year ---")
	for _, item := range year {
		fmt.Printf("#%-5d %9.9s %.55s\n",
			item.Number, item.User.Login, item.Title)
	}
}
