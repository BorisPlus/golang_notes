package main

// go run ./templator.go ./templator_app.go
func main() {
	//
	classificator_go := Template{}
	classificator_go.loadFromFile("./classificator.go", false)
	//
	notice := Template{}
	notice.loadFromFile("../templator/NOTICE.md", true)
	//
	substitutions := make(map[string]string)
	substitutions["classificator.go"] = tab_escaping(classificator_go.render())
	substitutions["notice"] = notice.render()
	//
	readme := Template{}
	readme.loadFromFile("./README.template.md", false)
	readme.substitutions = substitutions
	readme.renderToFile("./README.md")
}
