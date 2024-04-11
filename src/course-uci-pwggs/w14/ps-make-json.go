// Course 1 week 4
package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

func main() {

	mapStrStr := make(map[string]string)
	fmt.Println("Please input your name: ")
	br := bufio.NewReader(os.Stdin)
	bName, _, _ := br.ReadLine()
	name := string(bName)
	mapStrStr["name"] = name
	fmt.Println("Then input your address: ")
	br1 := bufio.NewReader(os.Stdin)
	bAddress, _, _ := br1.ReadLine()
	address := string(bAddress)
	mapStrStr["address"] = address
	bytesArray, err := json.Marshal(mapStrStr)
	if err != nil {
		fmt.Println("Marshalling failed")
	} else {
		fmt.Println("Bytes data: ")
		fmt.Println(bytesArray)
		fmt.Println("String data: ")
		fmt.Println(string(bytesArray))
	}
}
