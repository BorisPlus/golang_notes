package templator

import (
	"os"
	"strings"
)

type Template struct {
	Content       string
	Substitutions map[string]string
}

func (template *Template) LoadFromFile(filepath string, with_escaping bool) error {
	data, err := os.ReadFile(filepath)
	if err != nil {
		template.Content = ""
	}
	template.Content = string(data)
	if with_escaping {
		template.Content = escaping(template.Content)
	}
	return err
}

func (template *Template) Render() string {
	result := template.Content
	for k := range template.Substitutions {
		result = strings.Replace(result, "{{ "+k+" }}", template.Substitutions[k], -1)
	}
	return result
}

func escaping(content string) string {
	content = TabEscaping(content)
	content = strings.Replace(content, "\n```\n", "\n'''\n", -1)
	content = strings.Replace(content, "\n```text\n", "\n'''text\n", -1)
	content = strings.Replace(content, "\n```go\n", "\n'''go\n", -1)
	content = strings.Replace(content, "{{ ", "{"+string('\x02')+"{ ", -1)
	content = strings.Replace(content, " }}", " }"+string('\x02')+"}", -1)
	return content
}

func TabEscaping(content string) string {
	content = strings.Replace(content, "\t", "    ", -1)
	return content
}

func (template *Template) RenderToFile(filepath string) error {
	f, errCreate := os.Create(filepath)
	if errCreate != nil {
		return errCreate
	}
	result := template.Render()
	defer f.Close()

	_, errWrite := f.WriteString(result)
	if errWrite != nil {
		return errWrite
	}
	return nil
}
