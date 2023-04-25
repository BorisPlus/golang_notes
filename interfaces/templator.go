package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	template_file_name := "./README.template.md"
	template_file_data, err := ioutil.ReadFile(template_file_name)
	fmt.Print(template_file_data)
	if err != nil {
		fmt.Println("Err")
	}

	code_file_name := "./interface.go"
	code_file_data, err := ioutil.ReadFile(code_file_name)
	fmt.Print(code_file_data)

	a := strings.Replace(string(template_file_data), "content:"+code_file_name, string(code_file_data), -1)
	ioutil.WriteFile("README.md",
		[]byte(a),
		0644)
}
