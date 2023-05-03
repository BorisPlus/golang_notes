package main

// go run ./templator.go ./templator_app.go
func main() {
	//
	pointer_go := Template{}
	pointer_go.loadFromFile("./pointer.go", false)
	classificator_go := Template{}
	classificator_go.loadFromFile("./classificator.go", false)
	classificator_test_go := Template{}
	classificator_test_go.loadFromFile("./classificator_test.go", false)
	classificator_go_txt := Template{}
	classificator_go_txt.loadFromFile("./classificator.go.txt", false)
	//
	notice := Template{}
	notice.loadFromFile("../templator/NOTICE.md", true)
	//
	substitutions := make(map[string]string)
	substitutions["pointer.go"] = tab_escaping(pointer_go.render())
	substitutions["classificator.go"] = tab_escaping(classificator_go.render())
	substitutions["classificator_test.go"] = tab_escaping(classificator_test_go.render())
	substitutions["classificator.go.txt"] = classificator_go_txt.render()
	substitutions["notice"] = notice.render()
	//
	readme := Template{}
	readme.loadFromFile("./README.template.md", false)
	readme.substitutions = substitutions
	readme.renderToFile("./README.md")
}
