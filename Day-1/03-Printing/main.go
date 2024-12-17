package main

import (
	"fmt"
)

func main() {
	var userName string = "Ujjwal"
	userCaste := "Dahal"
	height := 3.5667
	fmt.Println("This is for Prnting But It Gives Line At End")

	fmt.Printf("Your name is %s & Caste is %s \n", userName, userCaste)
	fmt.Printf("Your Height is %.3f", height)
}
