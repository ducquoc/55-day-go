// Week 1
package main

import "fmt"

func main() {

	var lines int = 6
	PrintSailboat(lines, false)
}

// PrintSailboat - sailboat using ASCII art
func PrintSailboat(lines int, sparse bool) error {
	width := lines * 2
	for i := 0; i < width; i++ {
		if !sparse {
			for j := 0; j < lines; j++ {
				if i+j == lines {
					fmt.Print("/")
				} else if i == width/2 {
					fmt.Print("_")
				} else {
					fmt.Print(" ")
				}
			}
			for j := lines; j < width; j++ {
				if j > i && j-i == width/2-1 {
					fmt.Print("\\")
				} else if i == width/2 {
					fmt.Print("_")
				} else {
					fmt.Print(" ")
				}
			}
		}
		fmt.Println()
	}
	return nil
}
