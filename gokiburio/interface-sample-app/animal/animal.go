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

func (a *AnimalStatus) String() string {
	return fmt.Sprintf("Name: %s, Age: %s, Type: %s, Sount: %s", a.GetName(), a.GetAge(), a.GetType(), a.MakeSount())
}

func GetAnimalStatus(a Animal) {
	fmt.Println(a)
}
