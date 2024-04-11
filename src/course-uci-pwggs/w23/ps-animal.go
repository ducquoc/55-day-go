// Course 2 week 3
package main

import (
	"fmt"
)

type Animal struct {
	food       string
	locomotion string
	sound      string
}

func (a Animal) Eat() {
	fmt.Println(a.food)
}

func (a Animal) Move() {
	fmt.Println(a.locomotion)
}

func (a Animal) Speak() {
	fmt.Println(a.sound)
}

func main() {

	animalsActionsMap := map[string]Animal{
		"cow":   {"grass", "walk", "moo"},
		"bird":  {"worms", "fly", "peep"},
		"snake": {"mice", "slither", "hsss"},
	}
	fmt.Println("Please enter name [cow bird snake] and action [eat move speak] (Ctrl+C to stop):")
	for {
		fmt.Print("> ")
		var animal, action string
		fmt.Scan(&animal, &action)

		switch action {
		case "eat":
			animalsActionsMap[animal].Eat()
		case "move":
			animalsActionsMap[animal].Move()
		case "speak":
			animalsActionsMap[animal].Speak()
		}
	}
}
