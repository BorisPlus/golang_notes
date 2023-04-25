// package main

// import (
// 	"fmt"
// )

// func NamedFunction() func() int {
// 	anonymousFunction := func() int {
// 		return 0
// 	}
// 	return anonymousFunction
// }

// // func Sequence() func() int {
// // 	i := 0
// // 	AnonymousSequence := func() int {
// // 		i++
// // 		return i
// // 	}
// // 	return AnonymousSequence
// // }

// func main() {

// 	// exploit := func() int {
// 		// return 0
// 	// }

// 	// i := 1
// 	// exploit := i

// 	// fmt.Printf("              i\t%v\n", i)  // Address of exploit value of type `func() int`
// 	// fmt.Printf("             &i\t%p\n", &i) // Address of type `func() int` - function body `func() int { ... }`
// 	// fmt.Printf("        exploit\t%v\n", exploit) // Address of exploit value of type `func() int`nt`
// 	// fmt.Printf("       &exploit\t%p\n", &exploit) // Address of type `func() int` - function body `func() int { ... }`
// 	// fmt.Printf("    *(&exploit)\t%p\n", *(&exploit)) // Address of exploit value of type `func() int`
// 	// fmt.Println()

// 	// exploit = 3

// 	// fmt.Printf("              i\t%v\n", i)  // Address of exploit value of type `func() int`
// 	// fmt.Printf("             &i\t%p\n", &i) // Address of type `func() int` - function body `func() int { ... }`
// 	// fmt.Printf("        exploit\t%v\n", exploit) // Address of exploit value of type `func() int`nt`
// 	// fmt.Printf("       &exploit\t%p\n", &exploit) // Address of type `func() int` - function body `func() int { ... }`
// 	// fmt.Printf("    *(&exploit)\t%p\n", *(&exploit)) // Address of exploit value of type `func() int`
// 	// fmt.Println()

// 	// var exploitAddress *int
// 	// // var exploitAddress, anonAddress *func() int

// 	// fmt.Printf(" exploitAddress\t%v\n", exploitAddress)  // Address of exploit value
// 	// fmt.Printf("&exploitAddress\t%p\n", &exploitAddress) // Address of exploit function body - `func() int { ... }`

// 	// exploitAddress = &exploit

// 	// fmt.Printf(" exploitAddress\t%p\n", exploitAddress)  // Address of exploit value
// 	// fmt.Printf("&exploitAddress\t%p\n", &exploitAddress) // Address of exploit function body - `func() int { ... }`
	
// 	// *exploitAddress = i
	
// 	// fmt.Printf("              i\t%v\n", i)  // Address of exploit value of type `func() int`
// 	// fmt.Printf("             &i\t%p\n", &i) // Address of type `func() int` - function body `func() int { ... }`
// 	// fmt.Printf("        exploit\t%v\n", exploit) // Address of exploit value of type `func() int`nt`
// 	// fmt.Printf("       &exploit\t%p\n", &exploit) // Address of type `func() int` - function body `func() int { ... }`
// 	// fmt.Printf("    *(&exploit)\t%p\n", *(&exploit)) // Address of exploit value of type `func() int`
// 	// fmt.Println()

// 	// // fmt.Printf("    anonAddress\t%p\n", anonAddress)  // Address of exploit value
// 	// // fmt.Printf("   &anonAddress\t%p\n", &anonAddress) // Address of exploit function body - `func() int { ... }`

// 	// exploitAddress = &exploit

// 	// fmt.Printf(" exploitAddress\t%p\n", exploitAddress)  // Address of exploit value
// 	// fmt.Printf("&exploitAddress\t%p\n", &exploitAddress) // Address of exploit function body - `func() int { ... }`
// 	// fmt.Printf("*exploitAddress\t%p\n", *exploitAddress) // Address of exploit function body - `func() int { ... }`
// 	// // fmt.Printf("*exploitAddress\t%p\n", *(&exploit)) // Address of exploit function body - `func() int { ... }`

// 	// // fmt.Println()

// 	fmt.Printf("           NamedFunction\t%p\n", NamedFunction)  // Address of exploit value
// 	// fmt.Printf("          &NamedFunction\t%p\n", &NamedFunction) // Address of exploit value

// 	var function func() int
// 	fmt.Printf("                function\t%p\n", function)  // Address of exploit value
// 	fmt.Printf("               &function\t%p\n", &function) // Address of exploit value
// 	var functionResult int
// 	fmt.Printf("          functionResult\t%p\n", functionResult)  // Address of exploit value
// 	fmt.Printf("         &functionResult\t%p\n", &functionResult) // Address of exploit value
// 	var functionAddress *func() int
// 	fmt.Printf("         functionAddress\t%p\n", functionAddress)  // Address of exploit value
// 	fmt.Printf("        &functionAddress\t%p\n", &functionAddress) // Address of exploit value

// 	function = NamedFunction()
// 	fmt.Printf("                function\t%p\n", function)  // Address of exploit value
// 	fmt.Printf("               &function\t%p\n", &function) // Address of exploit value
// 	functionResult = function()

// 	var anonymousFunctionAddress *func() int
// 	anonymousFunctionAddress = &function

// 	fmt.Printf(" anonymousFunctionAddress\t%p\n", anonymousFunctionAddress)  // Address of exploit value
// 	fmt.Printf("&anonymousFunctionAddress\t%p\n", &anonymousFunctionAddress) // Address of exploit value

// 	// fmt.Println()
// 	// fmt.Println("--------------------------------------")
// 	// fmt.Println("       iterate = Sequence()")
// 	// fmt.Println()
// 	// // fmt.Printf("      &Sequence\t%p\n", &Sequence) // NotValid
// 	// fmt.Printf("       Sequence\t%v\n", Sequence) // Address of `func Sequence()`
// 	// fmt.Printf("        iterate\t%v\n", iterate)  // Address of current `AnonymousSequence` value
// 	// fmt.Printf("       &iterate\t%p\n", &iterate) // Address of `AnonymousSequence` function body - `func() int { ... }`

// 	// fmt.Println()
// 	// one = iterate()
// 	// fmt.Printf("            one\t%v\n", one)  // Iteation value
// 	// fmt.Printf("            one\t%p\n", &one) // Address of iteation value
// 	// fmt.Println()

// 	// fmt.Printf("       Sequence\t%p\n", Sequence) // Address of `func Sequence()`
// 	// fmt.Printf("        iterate\t%p\n", iterate)  // Address of current `AnonymousSequence` value
// 	// fmt.Printf("       &iterate\t%p\n", &iterate) // Address of `AnonymousSequence` function body - `func() int { ... }`

// 	// fmt.Println()
// 	// two = iterate()
// 	// fmt.Printf("            two\t%v\n", two)  // Iteation value
// 	// fmt.Printf("            two\t%p\n", &two) // Address of iteation value

// 	// fmt.Println()
// 	// fmt.Printf("       Sequence\t%p\n", Sequence) // Address of `func Sequence()`
// 	// fmt.Printf("        iterate\t%p\n", iterate)  // Address of current `AnonymousSequence` value
// 	// fmt.Printf("       &iterate\t%p\n", &iterate) // Address of `AnonymousSequence` function body - `func() int { ... }`

// 	// iterate = Sequence()

// 	// fmt.Println()
// 	// fmt.Println("--------------------------------------")
// 	// fmt.Println("       iterate = Sequence()")
// 	// fmt.Println()

// 	// iterate = exploit

// 	// fmt.Printf("       Sequence\t%p\n", Sequence) // Address of `func Sequence()`
// 	// fmt.Printf("        iterate\t%p\n", iterate)  // Address of current `AnonymousSequence` value
// 	// fmt.Printf("       &iterate\t%p\n", &iterate) // Address of `AnonymousSequence` function body - `func() int { ... }`

// 	// fmt.Println()
// 	// one = iterate()
// 	// fmt.Printf("            one\t%v\n", one)  // Iteation value
// 	// fmt.Printf("            one\t%p\n", &one) // Address of iteation value
// 	// fmt.Println()

// 	// fmt.Printf("       Sequence\t%p\n", Sequence) // Address of `func Sequence()`
// 	// fmt.Printf("        iterate\t%p\n", iterate)  // Address of current `AnonymousSequence` value
// 	// fmt.Printf("       &iterate\t%p\n", &iterate) // Address of intSeqAnon

// 	// fmt.Println()
// 	// two = iterate()
// 	// fmt.Printf("            two\t%v\n", two)  // Iteation value
// 	// fmt.Printf("            two\t%p\n", &two) // Address of iteation value

// 	// fmt.Println()
// 	// fmt.Printf("       Sequence\t%p\n", Sequence) // Address of `func Sequence()`
// 	// fmt.Printf("        iterate\t%p\n", iterate)  // Address of current `AnonymousSequence` value
// 	// fmt.Printf("       &iterate\t%p\n", &iterate) // Address of `AnonymousSequence` function body - `func() int { ... }`
// }
