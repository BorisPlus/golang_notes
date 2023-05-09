package main

// go run ./templator.go ./templator_app.go
func main() {
	//
	list_go := tTemplate{}
	list_go.loadFromFile("./list.go", false)
	list_stringer_go := tTemplate{}
	list_stringer_go.loadFromFile("./list_stringer.go", false)
	list_doc_txt := tTemplate{}
	list_doc_txt.loadFromFile("./list.doc.txt", false)
	list_test_go := tTemplate{}
	list_test_go.loadFromFile("./list_test.go", false)
	list_test_go_txt := tTemplate{}
	list_test_go_txt.loadFromFile("./list_test.go.txt", false)
	list_sort_test_go := tTemplate{}
	list_sort_test_go.loadFromFile("./list_sort_test.go", false)
	list_sort_test_go_txt := tTemplate{}
	list_sort_test_go_txt.loadFromFile("./list_sort_test.go.txt", false)
	//
	notice := tTemplate{}
	notice.loadFromFile("../templator/NOTICE.md", true)
	//
	substitutions := make(map[string]string)
	substitutions["list.go"] = tab_escaping(list_go.render())
	substitutions["list_stringer.go"] = tab_escaping(list_stringer_go.render())
	substitutions["list.doc.txt"] = list_doc_txt.render()
	substitutions["list_test.go"] = tab_escaping(list_test_go.render())
	substitutions["list_test.go.txt"] = list_test_go_txt.render()
	substitutions["list_sort_test.go"] = tab_escaping(list_sort_test_go.render())
	substitutions["list_sort_test.go.txt"] = list_sort_test_go_txt.render()
	substitutions["notice"] = notice.render()
	//
	readme := tTemplate{}
	readme.loadFromFile("./README.template.md", false)
	readme.substitutions = substitutions
	readme.renderToFile("./README.md")
}
