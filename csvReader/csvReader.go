package csvReader

import (
	"encoding/csv"
	"fmt"
	"os"
)

// CsvReader - Parse a csv for the quiz questions
func CsvReader(csvPath string) [][]string {
	recordFile, err := os.Open(csvPath)
	if err != nil {
		fmt.Println("An error encountered ::", err)
	}

	reader := csv.NewReader(recordFile)
	records, _ := reader.ReadAll()

	return records
}
