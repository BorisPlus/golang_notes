package main

import (
	"fmt"
	"math/rand"
)

func NamedFunction() (func() int, *func() int) {
	anonymousFunction := func() int {
		return rand.Intn(10)
	}
	return anonymousFunction, &anonymousFunction
}


var funtionCode string = "" +
`func NamedFunction() (func() int, *func() int) {
	anonymousFunction := func() int {
		return rand.Intn(10)
	}
	return anonymousFunction, &anonymousFunction
}
`

func main() {
	
	fmt.Println("```go")  // 
	fmt.Printf("%s",funtionCode)  // 
	fmt.Println("```")  // 
	fmt.Printf("\n")  // 

	fmt.Println("```go")  // 
	fmt.Printf("fmt.Printf(\"%%s\",NamedFunction)\n")  // 
	fmt.Println("```")  // 
	fmt.Println()  // 

	fmt.Printf("дает\n")  // 
	fmt.Println()  //  

	fmt.Println("```text")  // 
	fmt.Printf("%s\n", NamedFunction)  // 
	fmt.Println("```")  // 
	fmt.Println()  // 

	fmt.Printf("Оператор &NamedFunction недопустим, однако\n")  // 
	fmt.Println()  // 

	fmt.Println("```go")  // 
	fmt.Printf("fmt.Printf(\"%%p\",NamedFunction)\n")  // 
	fmt.Println("```")  // 
	fmt.Println()  // 

	fmt.Printf("дает\n")  // 
	fmt.Println()  //  

	fmt.Println("```text")  // 
	// fmt.Printf("%p\n", &NamedFunction)  // 
	fmt.Printf("%p\n", NamedFunction)  // 
	fmt.Println("```")  // 
	fmt.Println()  // 

	fmt.Printf("%p - это что такое? как это интерпритировать? если это адрес, то почему не 10 знаков после '0x'?\n", NamedFunction)  // 
	fmt.Printf("как это интерпритировать? если это адрес, то почему не 10 знаков после '0x'?\n")  // 
	fmt.Printf("если это адрес, то почему не 10 знаков после '0x'?\n")  // 

	// var function func() int
	// var functionAddress *func() int
	// fmt.Println("```go")  // 
	// fmt.Printf("var function func() int\n")  // 
	// fmt.Printf("var functionAddress *func() int\n")  // 
	// fmt.Println("```")  // 
	// fmt.Println()  // 
	// fmt.Println("```text")  // 
	// fmt.Printf("Different fmt.Printf:\n")  // 
	// fmt.Printf("          %%s of function\t%s\n", function)  // 
	// fmt.Printf("          %%p of function\t%p\n", function)  // 
	// fmt.Printf("         %%s of &function\t%s\n", &function) // 
	// fmt.Printf("         %%p of &function\t%p\n", &function) // 
	// fmt.Printf("   %%s of functionAddress\t%s\n", functionAddress)  // 
	// fmt.Printf("   %%p of functionAddress\t%p\n", functionAddress)  // 
	// fmt.Printf("  %%s of &functionAddress\t%s\n", &functionAddress) // 
	// fmt.Printf("  %%p of &functionAddress\t%p\n", &functionAddress) // 
	// fmt.Println("```")  // 
	// fmt.Println()  // 

	// function, functionAddress = NamedFunction()
	// fmt.Println("```go")  // 
	// fmt.Printf("function, functionAddress = NamedFunction()\n")  // 
	// fmt.Println("```")  // 
	// fmt.Println()  // 
	// fmt.Println("```text")  // 
	// fmt.Printf("Different fmt.Printf:\n")  // 
	// fmt.Printf("          %%s of function\t%s\n", function)  // 
	// fmt.Printf("          %%p of function\t%p\n", function)  // 
	// fmt.Printf("         %%s of &function\t%s\n", &function) // 
	// fmt.Printf("         %%p of &function\t%p\n", &function) // 
	// fmt.Printf("   %%s of functionAddress\t%s\n", functionAddress)  // 
	// fmt.Printf("   %%p of functionAddress\t%p\n", functionAddress)  // 
	// fmt.Printf("  %%s of &functionAddress\t%s\n", &functionAddress) // 
	// fmt.Printf("  %%p of &functionAddress\t%p\n", &functionAddress) // 
	// fmt.Printf("\n")  // 

	// function, functionAddress = NamedFunction()
	// fmt.Println("```go")  // 
	// fmt.Printf("function, functionAddress = NamedFunction()\n")  // 
	// fmt.Println("```")  // 
	// fmt.Println()  // 
	// fmt.Println("```text")  // 
	// fmt.Printf("Different fmt.Printf:\n")  // 
	// fmt.Printf("          %%s of function\t%s\n", function)  // 
	// fmt.Printf("          %%p of function\t%p\n", function)  // 
	// fmt.Printf("         %%s of &function\t%s\n", &function) // 
	// fmt.Printf("         %%p of &function\t%p\n", &function) // 
	// fmt.Printf("   %%s of functionAddress\t%s\n", functionAddress)  // 
	// fmt.Printf("   %%p of functionAddress\t%p\n", functionAddress)  // 
	// fmt.Printf("  %%s of &functionAddress\t%s\n", &functionAddress) // 
	// fmt.Printf("  %%p of &functionAddress\t%p\n", &functionAddress) // 
	// fmt.Printf("\n")  // 



	// fmt.Printf("   %%s of functionAddress\t%s\n", functionAddress)  // 
	// fmt.Printf("   %%p of functionAddress\t%p\n", functionAddress)  // 
	// fmt.Printf("  %%s of &functionAddress\t%s\n", &functionAddress) // 
	// fmt.Printf("  %%p of &functionAddress\t%p\n", &functionAddress) // 

	// var functionResult int
	// fmt.Printf("       %%s functionResult\t%s\n", functionResult)  // 
	// fmt.Printf("       %%p functionResult\t%p\n", functionResult)  // 
	// fmt.Printf("      %%s &functionResult\t%s\n", &functionResult)  // 
	// fmt.Printf("      %%p &functionResult\t%p\n", &functionResult) // 


	// function = NamedFunction()
	// fmt.Printf("                function\t%p\n", function)  // 
	// fmt.Printf("               &function\t%p\n", &function) // 
	// functionResult = function()
}
