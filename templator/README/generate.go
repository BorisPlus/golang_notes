package main

import (
	t "github.com/BorisPlus/golang_notes/templator"
)

// go run ./
func main() {
	//
	templator_go := t.Template{}
	templator_go.LoadFromFile("../templator.go", false)
	//
	templator_app_go := t.Template{}
	templator_app_go.LoadFromFile("../templator_app.go", false)
	//
	readme_template := t.Template{}
	readme_template.LoadFromFile("../README.template.md", true)
	//
	notice := t.Template{}
	notice.LoadFromFile("../NOTICE.md", true)
	//
	substitutions := make(map[string]string)
	substitutions["README.template.md"] = readme_template.Render()
	substitutions["templator.go"] = t.TabEscaping(templator_go.Render())
	substitutions["templator_app.go"] = t.TabEscaping(templator_app_go.Render())
	substitutions["notice"] = notice.Render()
	//
	readme := t.Template{}
	readme.LoadFromFile("../README.template.md", false)
	readme.Substitutions = substitutions
	readme.RenderToFile("../README.md")
}
