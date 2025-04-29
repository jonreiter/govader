package benchmark

import (
	"encoding/csv"
	"io"
	"os"
)

// LoadTextsFromCSV loads test texts from a CSV file
func LoadTextsFromCSV(filename string, textColumnIndex int) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var texts []string
	reader := csv.NewReader(file)

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}

		if textColumnIndex < len(record) {
			texts = append(texts, record[textColumnIndex])
		}
	}

	return texts, nil
}
