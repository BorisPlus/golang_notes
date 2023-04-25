package main

// go run templator_app.go templator.go
func main() {
	data_mapped := make(map[string]string)
	data_mapped["content:./interface.go"] = read("./interface.go")
	data_mapped["content:./interface_pointed_receiver.go"] = read("./interface_pointed_receiver.go")
	// 
	MakeReportFromTemplate("./README.template.md", data_mapped, "./README.md")
}
