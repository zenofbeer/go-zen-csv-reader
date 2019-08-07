package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	fmt.Println("Preparing to read file")
	// open file and create scanner on top of it.
	file, err := os.Open("./us-first-names-database-QueryResult (1).csv")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanLines)

	scanning := true

	for {
		success := scanner.Scan()
		if success == false {
			err = scanner.Err()
			if err == nil {
				log.Println("Scan completed and reached EOF.")
				scanning = success
			} else {
				log.Fatal()
			}
		} else {
			output := strings.Split(scanner.Text(), ",")
			fmt.Println(output[0], " - ", output[1])
		}
		if !scanning {
			break
		}
	}
}
