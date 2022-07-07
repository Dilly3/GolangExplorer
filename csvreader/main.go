package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func main() {
	students, err := readData("students0.csv")
	if err != nil {
		panic(err)
	}
	f, err := os.Create("students1.csv")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	writer := csv.NewWriter(f)
	defer writer.Flush()

	for _, student := range students {
		studentStruct := struct {
			id        string
			firstname string
			lastname  string
			email     string
		}{
			id:        student[0],
			firstname: student[1],
			lastname:  student[2],
			email:     student[3],
		}
		writer.Write(student)
		fmt.Println(studentStruct)

	}

}

func readData(filename string) ([][]string, error) {

	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	csv := csv.NewReader(file)

	// setting the delimiter to a comma
	csv.Comma = ','
	// setting # as comment
	csv.Comment = '#'
	csv.FieldsPerRecord = 4
	if _, err = csv.Read(); err != nil {
		return nil, err
	}
	var students [][]string
	if students, err = csv.ReadAll(); err != nil {
		return nil, err
	}
	return students, nil
}
