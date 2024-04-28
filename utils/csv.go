package utils

import (
	"encoding/csv"
	"log"
	"os"
)

func CSVtoJSON(path string) [][]string{
	csvFile, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer csvFile.Close()

	reader := csv.NewReader(csvFile)
	reader.FieldsPerRecord = -1

	csvData, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	return csvData
}
