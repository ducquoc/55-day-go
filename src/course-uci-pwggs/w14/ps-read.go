// Course 1 week 4
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

type Name struct {
	fname string
	lname string
}

func main() {

	path := "" // var path string
	fmt.Println("Please input file full path (such as /Users/dcqc/namesSlice.txt):")
	fmt.Scanln(&path)
	var namesSlice []Name
	f, err := os.Open(path)
	if err != nil {
		fmt.Println("[ERROR] panic(err): ", err.Error()) // https://youtrack.jetbrains.com/issue/GO-12179
	}
	defer f.Close() // https://youtrack.jetbrains.com/issue/GO-13454/Unresolved-reference-Close-for-os.File

	rd := bufio.NewReader(f)
	for {
		line, _, err := rd.ReadLine()

		if err != nil || io.EOF == err {
			break
		}
		splitNames := strings.SplitN(string(line), " ", 2)
		trimmedNames := Name{
			trimIfLongName(splitNames[0]),
			trimIfLongName(splitNames[1]),
		}
		namesSlice = append(namesSlice, trimmedNames)
	}
	for i := 0; i < len(namesSlice); i++ {
		fmt.Println(i + 1)
		fmt.Println("First Name: " + namesSlice[i].fname)
		fmt.Println("Last Name: " + namesSlice[i].lname)
	}
}

// Trim the string if the name is too long
func trimIfLongName(buffer string) string {

	if len(buffer) > 20 {
		return buffer[0:20] // string(...)
	} else {
		return buffer
	}
}
