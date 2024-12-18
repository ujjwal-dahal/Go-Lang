/*
you can explicitly use func functionName().

In a Go application, the main function is the entry point of the program. It doesn't take any
arguments and doesn't return anything. The execution of the program starts from the main
function.

. The function body is enclosed in curly braces {}.
. It's a Go convention that the opening curly brace { should be on the same line as the
function declaration.
*/

package main

import "fmt"

//Function with No Return Type

func displayName(userName string) {
	fmt.Printf("\n Your Name is %s", userName)
}

//Function with Return Type

func addNumber(a, b int) int {
	return a + b
}

//Function with Default return Type

func addNumber1(a, b int) (result int) {
	result = a + b
	return //Kai specify gareko chaina bhane result return huncha
}

func main() {

	displayName("Ujjwal")
	sum := addNumber(3, 4)
	fmt.Printf("\nSum is %d", sum)
	sum2 := addNumber1(10, 20)
	fmt.Printf("\n The Sum is %d", sum2)

}
