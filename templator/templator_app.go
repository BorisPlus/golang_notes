package main

// go run ./templator.go ./templator_app.go
func main() {
	// Карта подсток шаблона, заменяемых на конкретные значения, получаемые в результате исполнения шаблона
	data_mapped := make(map[string]string)
	data_mapped["var:notice"] = "П.С. Избегайте рекурсивной вложенности." // переменная
	data_mapped["content:./templator.go"] = read("./templator.go")
	data_mapped["content:./templator_app.go"] = read("./templator_app.go")
	//
	MakeReportFromTemplate("./README.template.md", data_mapped, "./README.md")

}
