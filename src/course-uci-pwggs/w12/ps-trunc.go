// Course 1 week 2
package main

import (
	"fmt"
)

func main() {

	var number int
	fmt.Println("Please enter a floating point number: ")
	fmt.Scanln(&number)
	fmt.Println("Truncated (rounded down) integer: ", number)
}
