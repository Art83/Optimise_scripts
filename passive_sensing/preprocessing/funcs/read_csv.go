package funcs

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

func ReadCsv(name string) ([]string, []string, []string) {
	var id []string
	var payload []string
	var date []string

	csvFile, err := os.Open(name)
	if err != nil {
		log.Fatal(err)
	}
	defer csvFile.Close()

	r := csv.NewReader(csvFile)
	fmt.Println("Reading the csv file")
	fmt.Println("---------------------")
	for {
		line, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Reading from the csv went wrong: %s", err)
		}
		id = append(id, line[0])
		date = append(date, line[2])
		payload = append(payload, line[1])
	}
	return id, payload, date
}
