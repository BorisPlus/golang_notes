package main

import (
	t "github.com/BorisPlus/golang_notes/templator"
)

// go run ./
func main() {
	//
	list_go := t.Template{}
	list_go.LoadFromFile("../list.go", false)
	list_stringer_go := t.Template{}
	list_stringer_go.LoadFromFile("../list_stringer.go", false)
	list_doc_txt := t.Template{}
	list_doc_txt.LoadFromFile("../list.doc.txt", false)
	list_test_go := t.Template{}
	list_test_go.LoadFromFile("../list_test.go", false)
	// 
	list_test_go_simple_txt := t.Template{}
	list_test_go_simple_txt.LoadFromFile("../list_test.go.simple.txt", false)
	list_test_go_complex_txt := t.Template{}
	list_test_go_complex_txt.LoadFromFile("../list_test.go.complex.txt", false)
	list_test_go_swap_txt := t.Template{}
	list_test_go_swap_txt.LoadFromFile("../list_test.go.swap.txt", false)
	list_test_go_sort_txt := t.Template{}
	list_test_go_sort_txt.LoadFromFile("../list_test.go.sort.txt", false)
	//
	notice := t.Template{}
	notice.LoadFromFile("../../templator/NOTICE.md", true)
	//
	substitutions := make(map[string]string)
	substitutions["list.go"] = t.TabEscaping(list_go.Render())
	substitutions["list_stringer.go"] = t.TabEscaping(list_stringer_go.Render())
	substitutions["list.doc.txt"] = list_doc_txt.Render()
	substitutions["list_test.go"] = t.TabEscaping(list_test_go.Render())
	// 
	substitutions["list_test.go.simple.txt"] = list_test_go_simple_txt.Render()
	substitutions["list_test.go.complex.txt"] = list_test_go_complex_txt.Render()
	substitutions["list_test.go.swap.txt"] = list_test_go_swap_txt.Render()
	substitutions["list_test.go.sort.txt"] = list_test_go_sort_txt.Render()
	substitutions["notice"] = notice.Render()
	//
	readme := t.Template{}
	readme.LoadFromFile("../README.template.md", false)
	readme.Substitutions = substitutions
	readme.RenderToFile("../README.md")
}
