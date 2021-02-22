package smartCSV

import (
	"encoding/csv"
	"os"
)

func readFile(filepath string) (result [][]string, isErr bool) {
	file, err := os.Open(filepath)
	if err != nil {
		return nil, true
	}

	defer file.Close()

	csvReader := csv.NewReader(file)
	var line []string
	var output [][]string

	for {
		line, err = csvReader.Read()
		if err != nil {
			return output, false
		}
		output = append(output, line)
	}
}
