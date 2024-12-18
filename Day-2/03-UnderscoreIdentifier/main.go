/*
Error Handling in Go

. Hey, In Go we Normally handle Errors as we do in other Programming Language, but
again Go Introduced a unique concept of_( underscore )
· In Go, the underscore (_) is used as a blank identifiğr. It serves as a write-only variable
that allows you to discard values returned by a function or to ignore specific values
when you are not interested in using them.
*/

package main

import "fmt"

//underscore ko kaam chai Function le kunai kura return gareko cha bhane teslai ignore garna lai ho

//error bhanne specific keyword nai cha
func divideNumber(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("denominator cannot be zero") //Errorf() le capitalize input lidaina
	}

	return a / b, nil
}

func main() {

	// result, err := divideNumber(10, 0) //abo err bhanne variable ma error aaucha ani yo compulsory use garnu parcha

	// if err != nil {
	// 	fmt.Println("Error Happen :", err)
	// }
	// fmt.Printf("\n Your Result is %f", result)

	result, _ := divideNumber(10, 0)

	fmt.Println("Your Result : ", result)

}
