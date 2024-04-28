package utils

import (
	"encoding/csv"
	"log"
	"os"
	"strconv"
)

type InputData struct {
	TotalIncome float64
	WHT float64
	Donation float64
}

func ExtractCSVData(path string) [][]string{
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

func CSVtoJSON(csvData [][]string) {
	var requestData []InputData

	for _, row := range csvData {
		totalIncome, err := strconv.ParseFloat(row[0], 64)
		if err != nil {
			log.Fatal(err)
		}

		wht, err := strconv.ParseFloat(row[1], 64)
		if err != nil {
			log.Fatal(err)
		}

		donation, err := strconv.ParseFloat(row[2], 64)
		if err != nil {
			log.Fatal(err)
		}

		requestData = append(requestData, InputData{
			TotalIncome: totalIncome,
			WHT: wht,
			Donation: donation,
		})
	}
}
