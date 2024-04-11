// Course 1 week 2
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	var str string
	fmt.Println("Please enter a string (starts with 'i', contains 'a', ends with 'n' ignorecase): ")
	//fmt.Scanln(&str)
	reader := bufio.NewReader(os.Stdin)
	str, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Unable to read STDIN, consider using Scan, e.g. Scanln(&str)")
	}
	str = strings.TrimSpace(str) // strings.ToLower(str) HasPrefix Contains HasSuffix
	if (str[0] == 'i' || strings.HasPrefix(str, "I")) &&
		strings.ContainsAny(str, "aA") &&
		(str[len(str)-1] == 'n' || strings.HasSuffix(str, "N")) {
		fmt.Print("Found!") // regexp.MatchString("[Ii].*[Aa].*[Nn]", input)
	} else {
		fmt.Print("Not Found!")
	}
}
