package main

import (
	"fmt"

	"github.com/MarkTBSS/029_Inheritance/living_instance"
)

func main() {
	living := living_instance.Living{}
	fmt.Println("Living : ")
	living.Walk()
	living.Eat()
	living.Breath()
	fmt.Println()

	devil := living_instance.Devil{}
	fmt.Println("Devil : ")
	devil.Walk()
	devil.Eat()
	devil.Breath()
	devil.Mutate()
	devil.Fly()
	fmt.Println()

	elf := living_instance.Elf{}
	fmt.Println("Elf : ")
	elf.Walk()
	elf.Eat()
	elf.Breath()
	elf.MagicSpell()
	fmt.Println()
}
