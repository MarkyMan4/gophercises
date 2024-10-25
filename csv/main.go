package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

type Record struct {
	col1 string
	col2 int
	col3 float64
}

func main() {
	file, _ := os.Open("data.csv")
	defer file.Close()

	reader := csv.NewReader(file)
	recs := []Record{}

	line, err := reader.Read()
	if err != nil {
		panic(err)
	}

	for line != nil {
		col1 := line[0]
		col2, intConvertErr := strconv.Atoi(line[1])
		col3, floatConvertErr := strconv.ParseFloat(line[2], 64)

		if intConvertErr != nil {
			panic(fmt.Errorf("failed to convert value %s to int", line[1]))
		}

		if floatConvertErr != nil {
			panic(fmt.Errorf("failed to convert value %s to float", line[2]))
		}

		lineRecord := Record{col1, col2, col3}
		recs = append(recs, lineRecord)

		line, err = reader.Read()
	}

	fmt.Println(recs)
}
