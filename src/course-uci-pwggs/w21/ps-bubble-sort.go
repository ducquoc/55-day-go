// Course 2 week 1
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	var numbers []int
	fmt.Println("Please enter integers in a line (to be sorted ascending order after newline):")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	for _, n := range strings.Fields(scanner.Text()) {
		num, err := strconv.Atoi(n)
		if err != nil {
			fmt.Println("[ERROR] panic(error)", err)
			return
		}
		numbers = append(numbers, num)
	}
	BubbleSort(numbers)
	fmt.Println(numbers)
}

// BubbleSort - sort by swapping
func BubbleSort(numbers []int) {
	for i := 0; i < len(numbers); i++ {
		for j := 0; j < len(numbers)-1-i; j++ {
			if numbers[j] > numbers[j+1] {
				Swap(numbers, j)
			}
		}
	}
}

// Swap - two numbers in a slice
func Swap(numbers []int, i int) {
	numbers[i], numbers[i+1] = numbers[i+1], numbers[i]
}
