package main

// go run ./
func main() {
	//
	templator_go := Template{}
	templator_go.loadFromFile("./templator.go", false)
	//
	templator_app_go := Template{}
	templator_app_go.loadFromFile("./templator_app.go", false)
	//
	readme_template := Template{}
	readme_template.loadFromFile("./README.template.md", true)
	//
	substitutions := make(map[string]string)
	substitutions["README.template.md"] = readme_template.render()
	substitutions["templator.go"] = tab_escaping(templator_go.render())
	substitutions["templator_app.go"] = tab_escaping(templator_app_go.render())
	substitutions["notice"] = "П.С. Избегайте рекурсивной вложенности."
	//
	readme := Template{}
	readme.loadFromFile("./README.template.md", false)
	readme.substitutions = substitutions
	readme.renderToFile("./README.md")
}
