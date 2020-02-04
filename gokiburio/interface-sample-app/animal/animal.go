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
	fmt.Println("Name: ", a.GetName())
	fmt.Println("Age: ", a.GetAge())
	fmt.Println("Type: ", a.GetType())
	fmt.Println("Sount: ", a.MakeSound())
}
