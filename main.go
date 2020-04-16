package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	fmt.Println("Preparing to read file")

	file, err := os.Open("./iris.data")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	reader := csv.NewReader(file)

	// reads unknown number of fields
	reader.FieldsPerRecord = -1

	rawData, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	// test to print out the data
	// ToDo: trim the trailing comma.
	for i := 0; i < len(rawData); i++ {
		println(strings.Join(rawData[i], ","))
	}
}
