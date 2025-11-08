package file

import (
	"encoding/csv"
	"fmt"
	"os"
)

func WriteToCsv() {
	fmt.Println("Started Writing")
	defer fmt.Println("Closing Writer")

	file, err := os.Create("./file/state.csv")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	m := map[string]string{"1": "One", "2": "Two", "3": "Three"}

	writer := csv.NewWriter(file)
	defer writer.Flush()

	writer.Write([]string{"Key", "Value"})
	for k, v := range m {
		in := []string{k, v}
		err := writer.Write(in)
		if err != nil {
			fmt.Println(err)
		}
	}

}

func ReadFromCsv() {
	fmt.Println("Started Reading")
	defer fmt.Println("Closing Reader")

	file, err := os.Open("./file/state.csv")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	m := map[string]string{}

	reader := csv.NewReader(file)

	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	for _, record := range records[1:] {
		m[record[0]] = record[1]
	}

	fmt.Println(m)
}
