package main

// go run ./templator.go ./templator_app.go ./interface.go ./printed.go
func main() {
	//
	interface_go := Template{}
	interface_go.loadFromFile("./interface.go", false)
	//
	interface_pointed_receiver_go := Template{}
	interface_pointed_receiver_go.loadFromFile("./interface_pointed_receiver.go", false)
	//
	readme_template := Template{}
	readme_template.loadFromFile("./README.template.md", true)
	//
	substitutions := make(map[string]string)
	substitutions["interface.go"] = tab_escaping(interface_go.render())
	substitutions["interface_pointed_receiver.go"] = tab_escaping(interface_pointed_receiver_go.render())
	substitutions["printed"] = printed()
	//
	readme := Template{"", substitutions}
	readme.loadFromFile("./README.template.md", false)
	readme.renderToFile("./README.md")
}
