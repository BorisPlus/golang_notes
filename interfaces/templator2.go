package main

import (
	"log"
	"os"
	"strings"
)

func read(file_name string) string {
	file_name_data, err := os.ReadFile(file_name)
	if err != nil {
		log.Fatal(err)
	}
	return string(file_name_data)
}

func MakeReportFromTemplate(template_file_name string, data map[string]string, result_file_name string) {

	template_file_data := read(template_file_name)

	result_file_data := template_file_data

	for k := range data {
		result_file_data = strings.Replace(result_file_data, k, data[k], -1)
	}

	result_file_data = strings.Replace(result_file_data, "\n\t\t\t\t", "\n                ", -1)
	result_file_data = strings.Replace(result_file_data, "\n\t\t\t", "\n            ", -1)
	result_file_data = strings.Replace(result_file_data, "\n\t\t", "\n        ", -1)
	result_file_data = strings.Replace(result_file_data, "\n\t", "\n    ", -1)

	f, errCreate := os.Create(result_file_name)
	if errCreate != nil {
		log.Fatal(errCreate)
	}

	defer f.Close()

	_, errWrite := f.WriteString(result_file_data)
	if errWrite != nil {
		log.Fatal(errWrite)
	}
}
