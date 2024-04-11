package living_instance

import "fmt"

type Living struct{}

func (l Living) Walk() {
	fmt.Println("Walking")
}

func (l Living) Eat() {
	fmt.Println("Eating")
}

func (l Living) Breath() {
	fmt.Println("Breathing")
}
