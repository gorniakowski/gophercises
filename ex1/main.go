package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"time"
)

func main() {
	filename := flag.String("csv", "problems.csv", "a csv file")
	seconds := flag.Int("t", 30, "number of seconds for timer ")
	flag.Parse()
	file, err := os.Open(*filename)
	if err != nil {
		fmt.Printf("Failed to open: %s", *filename)
		os.Exit(1)

	}
	r := csv.NewReader(file)
	counter, score := 0, 0
	var answer string
	fmt.Printf("You have got %d seconds. Press Enter to start", *seconds)
	fmt.Scanln()
	timer := time.NewTimer(time.Duration(*seconds) * time.Second)
	go func() {
		<-timer.C
		fmt.Printf("\n Time up! Your score is %d out of %d\n", score, counter)
		os.Exit(0)
	}()
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
		if answer == strings.TrimSpace(record[1]) {
			score++
		}
	}
	fmt.Printf("Your score is %d out of %d\n", score, counter)

}
