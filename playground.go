package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	// open the CSV - check for errors
	defer fmt.Println("hello!")
	csvfile, err := os.Open("data.csv")
	if err != nil {
		log.Fatalln("ERROR: ", err)
	}

	r := csv.NewReader(csvfile)

	//iterating through the records
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s, %s, %s\n", record[0], record[1], record[2])
	}

}
