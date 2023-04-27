# Интерфейсы

## Варианты

Оба варианта

* `value receiver`

```go
package main

import (
    "fmt"
)

type Adult interface {
    IsAdult() bool
    fmt.Stringer
}

type Person struct {
    age  int
    name string
}

func (p Person) IsAdult() bool {
    return p.age >= 18 
}

func (p Person) String() string {
    return fmt.Sprintf("%s is %d years old.", p.name , p.age)
}

func adultFilter(people []Adult) []Adult {
    adults := make([]Adult, 0)
    for _, p := range people {
        if p.IsAdult() {
            adults = append(adults, p)
        }
    }
    return adults
}

func _main() {
    people := []Adult{Person{15, "John"}, Person{18, "Joe"}, Person{45, "Mary"}}
    fmt.Println(adultFilter(people))
}

```

* `pointed receiver`

```go
package main

import (
    "fmt"
)

type Adult interface {
    IsAdult() bool
    fmt.Stringer
}

type Person struct {
    age  int
    name string
}

func (p *Person) IsAdult() bool {
    return p.age >= 18 
}

func (p *Person) String() string {
    return fmt.Sprintf("%s is %d years old.", p.name , p.age)
}

func adultFilter(people []Adult) []Adult {
    adults := make([]Adult, 0)
    for _, p := range people {
        if p.IsAdult() {
            adults = append(adults, p)
        }
    }
    return adults
}

func _main() {
    people := []Adult{&Person{15, "John"}, &Person{18, "Joe"}, &Person{45, "Mary"}}
    fmt.Println(adultFilter(people))
}

```

приведут к успешной фильтрации

```text
[Joe is 18 years old. Mary is 45 years old.]
```

## Поведение

```go
package main

import "fmt"

type MsgUserBalanceChanged struct {
    userID  string
    balance string
}

type MsgEventChanged struct {
    eventID string
}

func processMessage(msg interface{}) {

    switch message := msg.(type) {
    case MsgUserBalanceChanged:
        fmt.Printf("user %q balance was changed to %q\n", message.userID, message.balance)
    case MsgEventChanged:
        fmt.Printf("event %q was changed\n", message.eventID)
    default:
        fmt.Printf("unknown message: %q\n", message)
    }
}

func _dynamic() {
    processMessage(MsgUserBalanceChanged{"user-1", "1000"})
    processMessage(MsgEventChanged{"event-1"})
    processMessage("unknown")
}

// user "user-1" balance was changed to "1000"
// event "event-1" was changed
// unknown message: unknown

```
