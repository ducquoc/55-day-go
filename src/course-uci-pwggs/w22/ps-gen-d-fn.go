// Course 2 week 2
package main

import (
	"fmt"
)

func main() {

	var a, v0, s0, t float64
	fmt.Print("Please enter value of acceleration: ")
	fmt.Scanln(&a)

	fmt.Print("Please enter value of initial velocity: ")
	fmt.Scanln(&v0)

	fmt.Print("Please enter value of initial displacement: ")
	fmt.Scanln(&s0)

	fn := GenDisplaceFn(a, v0, s0)

	fmt.Print("Please enter value of the time: ")
	fmt.Scanln(&t)

	fmt.Println(fn(t))
}

// GenDisplaceFn - function for displacement
func GenDisplaceFn(a, v0, s0 float64) func(float64) float64 {

	return func(t float64) float64 {
		return ((a / 2) * t * t) + v0*t + s0
	}
}
