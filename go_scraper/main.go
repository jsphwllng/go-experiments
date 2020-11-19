package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/gocolly/colly"
)

func scraper() (title string) {
	c := colly.NewCollector()
	c.OnHTML("h1", func(e *colly.HTMLElement) {
		title = e.Text
	})
	c.Visit("https://en.wikipedia.org/wiki/Special:Random")
	return title
}

func checkError(message string, err error) {
	if err != nil {
		log.Fatal(message, err)
	}
}

func csvWriter(title string) {
	time := time.Now().Format("Mon Jan _2 15:04:05 2006")
	//unusual OpenFile arguments here
	file, err := os.OpenFile("titles.csv", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	checkError("Cannot open file", err)
	writer := csv.NewWriter(file)
	err = writer.Write([]string{title, time})
	checkError("Cannot write to file", err)
	writer.Flush()
}

func main() {
	fmt.Println("beginning the scrape now:")
	title := scraper()
	for i := 0; i < 100; i++ {
		if string([]rune(title)[0]) != "R" || title == "" {
			title = scraper()
			i--
		} else {
			title = strings.Replace(title, "\"", "", -1)
			csvWriter(title)
			fmt.Printf("We now have %d titles\n", i+1)
			title = scraper()
		}
	}
}
