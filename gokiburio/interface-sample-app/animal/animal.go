package animal

import "fmt"

type AnimalStatus struct {
	Name string
	Age  int
	Type string
}

type Animal interface {
	MakeSound() string
	GetName() string
	GetAge() int
	GetType() string
}

func GetAnimalStatus(a Animal) {
	fmt.Sprintln("Name: %s", a.GetName())
	fmt.Sprintln("Age: %d", a.GetAge())
	fmt.Sprintln("Type: %s", a.GetType())
	fmt.Sprintln("Sount: %s", a.MakeSound())
}
