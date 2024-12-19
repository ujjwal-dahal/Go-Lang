/*
 => Array in Golang

Arrays in Go provides a simple and efficient way to work with a fixed collection of
elements.
Suppose you want to store a List of fruits in the array.

#Array Declaration:

Arrays are declared using the syntax var name [size]type, where name is
the array variable, size is the number of elements, and type is the data type of
the elements.


Arrays have a fixed size, meaning you must specify the number of elements the
array can hold at the time of declaration.

Once the array is created, its size cannot be changed.

Example: var numbers [5]int declares an integer array with five elements.


#Array Initialization:
Arrays can be initialized at the time of declaration or later in the program.
Example: var numbers [5]int creates an array of integers with zero values.

You can also initialize it with specific values: var numbers = [5]int{1, 2,
3, 4, 5}.

*/

package main

import "fmt"

func main() {
	// var number = [5]int{1, 2, 3, 4, 5}
	// // fmt.Printf("Array Elements are : %v and Its Type is %T", number, number)

	// fmt.Printf("First Element is %d", number[0])

	// var studentDetails = [3]string{"Ujjwal", "Hari", "Shyam"}
	// fmt.Printf("Your Name is %s", studentDetails[0])

	// var personName [5]string
	// personName[0] = "Ujjwal"
	// personName[1] = "Hari"

	// // fmt.Printf("Persons Name are : %v", personName)

	// //Finding Length of Array
	// fmt.Printf("Length of Array is %d", len(personName))

	/*
		In Go, when you declare an array or a slice, the elements are initialized to their zero
		values.
		The zero value is the default value assigned to variables of a certain type when
		no explicit value is provided.
		For numeric types (int, float, etc.), the zero value is 0.
		For strings, it is an empty string("").
		For boolean types, it is false, and
		for pointers or complex types, it is nil.


		-> Matlab hamile Array initialize garda esle Default value haldincha Go Language ma
		int type ko cha bhane 0 haldincha
		string type ko cha bhane empty string haldincha
		boolean type ko cha bhane false haldincha
		pointers type ko cha bhane nil haldincha
	*/

	var number = [5]int{1, 2, 3}
	fmt.Println(number) //Abo Hamile Index 0 1 2 ma value haldeko chau but 3 4 ma chaina so default value jancha i.e 0

}
