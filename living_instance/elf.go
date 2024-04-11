package living_instance

import "fmt"

type Elf struct {
	Living
}

func (e Elf) MagicSpell() {
	fmt.Println("Casting a magic spell")
}
