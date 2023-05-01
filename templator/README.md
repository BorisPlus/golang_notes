# Шаблонизатор отчетов

Этот отчет был составлен с задействованием

* шаблона [README.template.md](./README.template.md)

<details>
<summary>см. шаблон</summary>

```text
# Шаблонизатор отчетов

Этот отчет был составлен с задействованием

* шаблона [README.template.md](./README.template.md)

<details>
<summary>см. шаблон</summary>

'''text
{{ README.template.md }}
'''

</details>

* структуры

'''go
{{ templator.go }}
'''

* приложения

'''go
{{ templator_app.go }}
'''

* посредством вызова

```shell
go run ./templator.go ./templator_app.go 
'''

{{ notice }}

```

</details>

* структуры

```go
package main

import (
    "os"
    "strings"
)

type Template struct {
    content       string
    substitutions map[string]string
}

func (template *Template) loadFromFile(filepath string, with_escaping bool) error {
    data, err := os.ReadFile(filepath)
    if err != nil {
        template.content = ""
    }
    template.content = string(data)
    if with_escaping {
        template.content = escaping(template.content)
    }
    return err
}

func (template *Template) render() string {
    result := template.content
    for k := range template.substitutions {
        result = strings.Replace(result, "{{ "+k+" }}", template.substitutions[k], -1)
    }
    return result
}

func escaping(content string) string {
    content = tab_escaping(content)
    content = strings.Replace(content, "\n```\n", "\n'''\n", -1)
    content = strings.Replace(content, "\n```text\n", "\n'''text\n", -1)
    content = strings.Replace(content, "\n```go\n", "\n'''go\n", -1)
    content = strings.Replace(content, "{{ ", "{"+string('\x02')+"{ ", -1)
    content = strings.Replace(content, " }}", " }"+string('\x02')+"}", -1)
    return content
}

func tab_escaping(content string) string {
    content = strings.Replace(content, "\t", "    ", -1)
    return content
}

func (template *Template) renderToFile(filepath string) error {
    f, errCreate := os.Create(filepath)
    if errCreate != nil {
        return errCreate
    }
    result := template.render()
    defer f.Close()

    _, errWrite := f.WriteString(result)
    if errWrite != nil {
        return errWrite
    }
    return nil
}

```

* приложения

```go
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
    notice := Template{}
    notice.loadFromFile("./NOTICE.md", true)
    //
    substitutions := make(map[string]string)
    substitutions["README.template.md"] = readme_template.render()
    substitutions["templator.go"] = tab_escaping(templator_go.render())
    substitutions["templator_app.go"] = tab_escaping(templator_app_go.render())
    substitutions["notice"] = notice.render()
    //
    readme := Template{}
    readme.loadFromFile("./README.template.md", false)
    readme.substitutions = substitutions
    readme.renderToFile("./README.md")
}

```

* посредством вызова

```shell
go run ./templator.go ./templator_app.go 
```

> ```text
> Данный документ составлен с использованием 
> разработанного [шаблонизатора](https://github.com/BorisPlus/golang_notes/tree/master/templator). 
> При его использовании избегайте рекурсивной вложенности.
> ```
