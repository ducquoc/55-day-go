// Course 1 week 3
package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {

	slice := make([]int, 3)
	var inStr string
	fmt.Println("Init integer slice of length 3:", slice)
	for true {
		fmt.Println("Please enter an integer to be appended to slice then sort and print (X to exit): ")
		fmt.Scanln(&inStr)
		if inStr == "X" {
			break
		}
		numToBeAppended, err := strconv.Atoi(inStr)
		if err != nil {
			fmt.Println("Wrong input")
			continue
		}
		slice = append(slice, numToBeAppended)
		sort.Ints(slice[:]) // slices.Sort()
		fmt.Println(slice)
	}
}
