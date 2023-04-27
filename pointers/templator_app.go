package main

import (
	"fmt"
)

// go run ./
func main() {
	//
	pointer_go := Template{}
	pointer_go.loadFromFile("./pointer.go", false)
	//
	readme_template := Template{}
	readme_template.loadFromFile("./README.template.md", true)
	//
	substitutions := make(map[string]string)
	substitutions["pointer.go"] = tab_escaping(pointer_go.render())
	substitutions["sFunctionWithName"] = fmt.Sprintf("%s", FunctionWithName)
	substitutions["pFunctionWithName"] = fmt.Sprintf("%p", FunctionWithName)
	//
	readme := Template{"", substitutions}
	readme.loadFromFile("./README.template.md", false)
	readme.renderToFile("./README.md")
}
