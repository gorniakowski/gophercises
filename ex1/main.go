package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {

	filename := flag.String("csv", "problems.csv", "a csv file")
	flag.Parse()
	file, err := os.Open(*filename)
	if err != nil {
		fmt.Printf("Failed to open: %s", *filename)
		os.Exit(1)

	}
	r := csv.NewReader(file)
	counter, score := 0, 0
	var answer string
	for {
		counter++
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Problem #%d: %s=", counter, record[0])
		_, err = fmt.Scanln(&answer)
		if err != nil {
			fmt.Println(err)
			break
		}
		if answer == record[1] {
			score++
		}
	}
	fmt.Printf("Your score is %d out of %d", score, counter)

}
