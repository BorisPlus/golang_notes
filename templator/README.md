# Шаблонизатор отчетов

Этот отчет был составлен с задействованием функции

```go
package main

import (
    "log"
    "os"
    "strings"
)

func read(file_name string) string {
    file_name_data, err := os.ReadFile(file_name)
    if err != nil {
        log.Fatal(err)
    }
    // TODO: нужна опция
    // file_name_data_str := string(file_name_data)
    // file_name_data_str = strings.Replace(file_name_data_str, "```", "'''", -1)
    return string(file_name_data)
}

func MakeReportFromTemplate(template_file_name string, data map[string]string, result_file_name string) {

    template_file_data := read(template_file_name)

    result_file_data := template_file_data

    for k := range data {
        result_file_data = strings.Replace(result_file_data, k, data[k], -1)
    }

    // TODO: нужно изящнее 
    result_file_data = strings.Replace(result_file_data, "\n\t\t\t\t", "\n                ", -1)
    result_file_data = strings.Replace(result_file_data, "\n\t\t\t", "\n            ", -1)
    result_file_data = strings.Replace(result_file_data, "\n\t\t", "\n        ", -1)
    result_file_data = strings.Replace(result_file_data, "\n\t", "\n    ", -1)

    f, errCreate := os.Create(result_file_name)
    if errCreate != nil {
        log.Fatal(errCreate)
    }

    defer f.Close()

    _, errWrite := f.WriteString(result_file_data)
    if errWrite != nil {
        log.Fatal(errWrite)
    }
}

```

посредством вызова

```shell
go run ./templator.go ./templator_app.go 
```

приложения

```go
package main

// go run ./templator.go ./templator_app.go
func main() {
    // Карта подсток шаблона, заменяемых на конкретные значения, получаемые в результате исполнения шаблона
    data_mapped := make(map[string]string)
    data_mapped["var:notice"] = "П.С. Избегайте рекурсивной вложенности." // переменная
    data_mapped["content:./templator.go"] = read("./templator.go")
    data_mapped["content:./templator_app.go"] = read("./templator_app.go")
    //
    MakeReportFromTemplate("./README.template.md", data_mapped, "./README.md")

}

```

и шаблона [README.template.md](./README.template.md)

П.С. Избегайте рекурсивной вложенности.
