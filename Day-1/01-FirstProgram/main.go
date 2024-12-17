package main

import (
	"first_project/adding_data"
	displaystring "first_project/displayString"
	"fmt"
)

func main() {
	fmt.Println("Welcome to Golang Program")

	// Correct syntax for declaring a variable in Go
	result := adding_data.AddData(1,2)

	fmt.Printf("The result is: %d\n", result)

	//Display User Name
	 userName := "Ujjwal"
	displaystring.DisplayString(userName)
}
