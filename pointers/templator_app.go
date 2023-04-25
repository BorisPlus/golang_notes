package main

import (
	"fmt"
	"strings"
)


// go run ./pointer.go ./templator.go ./templator_app.go 
func main() {
	data_mapped := make(map[string]string)
	// 
	data_mapped["content:./pointer.go"] = strings.Replace(read("./pointer.go"), "```", "'''", -1)
	data_mapped["var:sFunctionWithName"] = fmt.Sprintf("%s", FunctionWithName)
	data_mapped["var:pFunctionWithName"] = fmt.Sprintf("%p", FunctionWithName)
	// 
	MakeReportFromTemplate("./README.template.md", data_mapped, "./README.md")

}
