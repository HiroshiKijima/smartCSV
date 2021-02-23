package smartCSV

import (
	"encoding/csv"
	"os"
)

func readFile(filepath string, doesSkipFirstLine bool) (result [][]string, isErr bool) {
	file, err := os.Open(filepath)
	if err != nil {
		return nil, true
	}

	defer file.Close()

	csvReader := csv.NewReader(file)
	var line []string
	var output [][]string
	isFirstLine := true

	for {
		line, err = csvReader.Read()
		if err != nil {
			return output, false
		}
		if doesSkipFirstLine && isFirstLine {
			isFirstLine = false
			continue
		}
		output = append(output, line)
	}
}
