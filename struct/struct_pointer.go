package main

import "fmt"

// import (
// 	"fmt"
// )

type StructuredValue struct {
	field int
}

type StructuredPointer struct {
	field *int
}

func main() {
	value := 1
	// Address of value 0xc0000180c8

	fmt.Printf(`Address of value %p`, &value)
	fmt.Println()

	sv := StructuredValue{value}
	fmt.Printf(`Address of structure object %p, 
	address of structure object first field %v contains pointer %v \ %v with value %v`,
		&sv, &(sv).field, &(sv.field), &sv.field, sv.field)
	fmt.Println()
	// Address of structure object 0xc0000180d0,
	// address of structure object first field 0xc0000180d0

	sp := StructuredPointer{&value}
	fmt.Printf(`Address of structure object %p, 
	address of structure object first field %v contains pointer %v \ %v with value %v`,
		&sp, &(sp).field, &(sp.field), &sp.field, sp.field)
	fmt.Println()
	// Address of structure object 0xc000012030,
	// address of structure object first field 0xc0000180c8

}
