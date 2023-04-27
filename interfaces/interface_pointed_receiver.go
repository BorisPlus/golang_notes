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
