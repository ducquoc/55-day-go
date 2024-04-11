// Course 2 week 4
package main

import (
	"fmt"
)

type Animal interface {
	Eat()
	Move()
	Speak()
}

type Cow struct{}
type Bird struct{}
type Snake struct{}

func (c *Cow) Eat() {
	fmt.Println("grass")
}

func (c *Cow) Move() {
	fmt.Println("walk")
}

func (c *Cow) Speak() {
	fmt.Println("moo")
}

func (b *Bird) Eat() {
	fmt.Println("worms")
}

func (b *Bird) Move() {
	fmt.Println("fly")
}

func (b *Bird) Speak() {
	fmt.Println("peep")
}

func (s *Snake) Eat() {
	fmt.Println("mice")
}

func (s *Snake) Move() {
	fmt.Println("slither")
}

func (s *Snake) Speak() {
	fmt.Println("hsss")
}

func CreateAnimal(a string) Animal {
	switch a {
	case "cow":
		return &Cow{}
	case "bird":
		return &Bird{}
	case "snake":
		return &Snake{}
	default:
		fmt.Errorf("error: invalid animal (available types: cow, bird, and snake).")
		return nil
	}
}

func QueryAnimal(a Animal, q string) {
	if a == nil {
		fmt.Println("error not found animal by name (valid names are created from 'newanimal' command).")
		return
	}
	switch q {
	case "move":
		a.Move()
	case "eat":
		a.Eat()
	case "speak":
		a.Speak()
	default:
		fmt.Printf("error: invalid operation (available types: move, eat, and speak).")
	}
}

func main() {

	animals := make(map[string]Animal)
	fmt.Println("Please enter 1 of [newanimal query] and 1 name and 1 type [cow bird snake] (Ctrl+C to stop):")
	for {
		fmt.Print("> ")
		var request, name, option string
		fmt.Scan(&request, &name, &option)

		switch request {
		case "newanimal":
			a := CreateAnimal(option)
			if a == nil {
				fmt.Errorf("error creating animal.")
			}
			animals[name] = a
			fmt.Println("Created it!")
		case "query":
			a, ok := animals[name]
			if !ok {
				fmt.Errorf("error: animal not found!")
			}
			QueryAnimal(a, option)
		default:
			fmt.Errorf("error: invalid request (try 'newanimal' or 'query').")
		}
	}
}
