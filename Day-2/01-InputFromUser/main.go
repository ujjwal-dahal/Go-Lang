/*
In this example:

1. We declare a variable name of type string to store the user's input.
2. We use fmt. Println to display a prompt asking the user to enter their name.
3. We use fmt.Scan(&name) to read the user's input and store it in the name variable.
The & symbol is used to get the memory address of the variable for Scan to populate.
4. Finally, we use fmt. Printf to print a greeting with the user's name.

. Keep in mind that fmt. Scan reads until the first whitespace character, so if you want to
read a whole line, you might want to use bufio package's NewScanner or
ReadString functions for more complex scenarios.

-> Scan() le WhiteSpace samma matra Input lincha
Example : If User Enter his name : Ujjwal Dahal
then Variable ma Ujjwal matra store huncha

-> So eslai Fix garna we use "bufio" package

*/

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	fmt.Println("Taking Input From User")
	fmt.Println("Enter Your Name : ")
	// var userName string
	// fmt.Scan(&userName)
	// fmt.Printf("\nYour Name is %s", userName)

	//reader banaune jasko kaam read garne ho input
	reader := bufio.NewReader(os.Stdin)

	//abo reader le kaile samma read garne bhanne
	userName, _ := reader.ReadString('\n') //next line na aaunjel read garcha
	fmt.Printf("\n Your Name is %s", userName)

}

/*
. bufio.NewReader(os.Stdin) creates a new buffered reader that reads from the
standard input (os. Stdin).
. reader. ReadString('\n') reads a line from the input until it encounters a newline
character ('\n'). This allows the program to read the entire line of input, including
spaces.
. The input is stored in the name variable, and any potential errors are handled.

· A buffered reader is a type of reader that reads data from an underlying source, such as
a file or standard input (keyboard), and stores that data in a buffer. The buffer is a
temporary storage area in memory. Buffered readers are commonly used to improve the
efficiency of input operations.
*/
