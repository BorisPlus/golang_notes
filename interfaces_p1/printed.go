package main

import (
	"fmt"
)

func printed() string{
	people := []Adult{Person{15, "John"}, Person{18, "Joe"}, Person{45, "Mary"}}
	return fmt.Sprint(adultFilter(people))
}
