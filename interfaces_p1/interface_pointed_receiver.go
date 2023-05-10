package main

import (
	"fmt"
)

type PAdult interface {
	IsPAdult() bool
	fmt.Stringer
}

type PPerson struct {
	age  int
	name string
}

func (p *PPerson) IsPAdult() bool {
	return p.age >= 18 
}

func (p *PPerson) String() string {
	return fmt.Sprintf("%s is %d years old.", p.name , p.age)
}

func adultPFilter(people []Adult) []Adult {
	adults := make([]Adult, 0)
	for _, p := range people {
		if p.IsAdult() {
			adults = append(adults, p)
		}
	}
	return adults
}

func Pmain() {
	people := []Adult{&Person{15, "John"}, &Person{18, "Joe"}, &Person{45, "Mary"}}
	fmt.Println(adultFilter(people))
}
