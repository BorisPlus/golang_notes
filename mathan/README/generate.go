package main

import (
	t "github.com/BorisPlus/golang_notes/templator"
)

func main() {
	//
	pointer_go := t.Template{}
	pointer_go.LoadFromFile("../pointer/pointer.go", false)
	classificator_go := t.Template{}
	classificator_go.LoadFromFile("../classificator/classificator.go", false)
	classificator_test_go := t.Template{}
	classificator_test_go.LoadFromFile("../classificator/classificator_test.go", false)
	classificator_go_txt := t.Template{}
	classificator_go_txt.LoadFromFile("../classificator/classificator.go.txt", false)
	clusterizator_go := t.Template{}
	clusterizator_go.LoadFromFile("../clusterizator/clusterizator.go", false)
	clusterizator_test_go := t.Template{}
	clusterizator_test_go.LoadFromFile("../clusterizator/clusterizator_test.go", false)
	clusterizator_go_txt := t.Template{}
	clusterizator_go_txt.LoadFromFile("../clusterizator/clusterizator.go.txt", false)
	//
	notice := t.Template{}
	notice.LoadFromFile("../templator/NOTICE.md", true)
	//
	substitutions := make(map[string]string)
	substitutions["pointer.go"] = t.TabEscaping(pointer_go.Render())
	substitutions["classificator.go"] = t.TabEscaping(classificator_go.Render())
	substitutions["classificator_test.go"] = t.TabEscaping(classificator_test_go.Render())
	substitutions["classificator.go.txt"] = classificator_go_txt.Render()
	substitutions["clusterizator.go"] = t.TabEscaping(clusterizator_go.Render())
	substitutions["clusterizator_test.go"] = t.TabEscaping(clusterizator_test_go.Render())
	substitutions["clusterizator.go.txt"] = t.TabEscaping(clusterizator_go_txt.Render())
	substitutions["notice"] = notice.Render()
	//
	readme := t.Template{}
	readme.LoadFromFile("../README.template.md", false)
	readme.Substitutions = substitutions
	readme.RenderToFile("../README.md")
}
