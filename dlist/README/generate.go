package main

import (
	t "github.com/BorisPlus/golang_notes/templator"
)

// go run ./
func main() {
	//
	dlist_go := t.Template{}
	dlist_go.LoadFromFile("../dlist.go", false)
	dlist_stringer_go := t.Template{}
	dlist_stringer_go.LoadFromFile("../dlist_stringer.go", false)
	dlist_doc_txt := t.Template{}
	dlist_doc_txt.LoadFromFile("../dlist.doc.txt", false)
	dlist_test_go := t.Template{}
	dlist_test_go.LoadFromFile("../dlist_test.go", false)
	// 
	dlist_test_go_simple_txt := t.Template{}
	dlist_test_go_simple_txt.LoadFromFile("../dlist_test.go.simple.txt", false)
	dlist_test_go_complex_txt := t.Template{}
	dlist_test_go_complex_txt.LoadFromFile("../dlist_test.go.complex.txt", false)
	dlist_test_go_swap_txt := t.Template{}
	dlist_test_go_swap_txt.LoadFromFile("../dlist_test.go.swap.txt", false)
	dlist_test_go_sort_txt := t.Template{}
	dlist_test_go_sort_txt.LoadFromFile("../dlist_test.go.sort.txt", false)
	//
	notice := t.Template{}
	notice.LoadFromFile("../../templator/NOTICE.md", true)
	//
	substitutions := make(map[string]string)
	substitutions["dlist.go"] = t.TabEscaping(dlist_go.Render())
	substitutions["dlist_stringer.go"] = t.TabEscaping(dlist_stringer_go.Render())
	substitutions["dlist.doc.txt"] = dlist_doc_txt.Render()
	substitutions["dlist_test.go"] = t.TabEscaping(dlist_test_go.Render())
	// 
	substitutions["dlist_test.go.simple.txt"] = dlist_test_go_simple_txt.Render()
	substitutions["dlist_test.go.complex.txt"] = dlist_test_go_complex_txt.Render()
	substitutions["dlist_test.go.swap.txt"] = dlist_test_go_swap_txt.Render()
	substitutions["dlist_test.go.sort.txt"] = dlist_test_go_sort_txt.Render()
	substitutions["notice"] = notice.Render()
	//
	readme := t.Template{}
	readme.LoadFromFile("../README.template.md", false)
	readme.Substitutions = substitutions
	readme.RenderToFile("../README.md")
}
