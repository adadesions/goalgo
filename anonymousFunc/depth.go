package main

import (
	"fmt"
)

func Top(data map[string][]string) {
	var search func(target string) bool
	var readSubjects func(items []string)

	search = func(target string) bool {
		fmt.Printf("Finding: %s\n", target)
		for key, item := range data {
			if key == target {
				fmt.Printf("Student: %v\n", key)
				readSubjects(item)
				return true
			}
		}

		return false
	}

	readSubjects = func(items []string) {
		for idx, item := range items {
			fmt.Printf("%d:\t%s\n", idx, item)
		}
	}

	isFound := search("Ada")
	fmt.Printf("isFound: %v\n", isFound)
}

func AnyDict(data map[string]map[string]interface{}) {
	for k, v := range data {
		fmt.Printf("%s:\t%v\n", k, v)
	}
}

type BookInfo struct {
	Name   string
	Author string
	Price  float32
}

type Review struct {
	ReviewerName string
	Content      string
}

type GoodBook struct {
	BookInfo
	Rating  int
	Reviews []Review
}

func (gb *GoodBook) Print() {
	fmt.Println("====== Good Book Info ======")
	fmt.Printf("Book Name:\t %s\n", gb.Name)
	fmt.Printf("Author:\t %s\n", gb.Author)
	fmt.Printf("Price:\t $ %f\n", gb.Price)
	fmt.Println("====== Rating ======")
	if gb.Rating > 50 {
		fmt.Printf("Rating (%d):\t Very Good\n", gb.Rating)
	} else {
		fmt.Printf("Rating (%d):\t Bad\n", gb.Rating)
	}
	fmt.Println("====== Reviews ======")
	for idx, review := range gb.Reviews {
		fmt.Printf("No.%d Customer Name: %s\n", idx, review.ReviewerName)
		fmt.Printf("\"%s\"\n", review.Content)
	fmt.Println("====================")
	}
}

func main() {
	var data = map[string][]string{
		"Ada":   {"Dynamic Programming", "Discrete Math"},
		"Ping":  {"Math POSN", "Master Class"},
		"Peter": {"How to start a Spiderman career", "How to become a hero"},
	}

	Top(data)

	var dataAny map[string]map[string]interface{}
	dataAny = map[string]map[string]interface{}{
		"Books": {
			"name":  "Go Lang",
			"price": 123.4,
		},
	}

	AnyDict(dataAny)

	goodBook := GoodBook{
		BookInfo: BookInfo{
			Name:   "Intro to Adaverse",
			Author: "Ada De Sions",
			Price:  399.99,
		},
		Rating: 100,
		Reviews: []Review{
			{ReviewerName: "Peter Parker", Content: "I really love your book Ada, nice work keep going."},
			{ReviewerName: "Capt. America", Content: "Nice work keep going!!!"},
			{ReviewerName: "T. Stark", Content: "The best book ever I've read!"},
		},
	}


	goodBook.Print()
}
