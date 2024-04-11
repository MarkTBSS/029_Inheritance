package living_instance

import "fmt"

type Devil struct {
	Living
}

func (d Devil) Mutate() {
	fmt.Println("Mutating")
}

func (d Devil) Fly() {
	fmt.Println("Flying")
}
