package car

import (
	"fmt"
	"strconv"
)

type CarValue struct {
	Name      string
	MaxSpeed  int
	TireCount int
}

type Car interface {
	Name() string
	MaxSpeed() int
	TireCount() int
}

func PrintCar(c Car) {
	fmt.Println("Name=" + c.Name())
	fmt.Println("MaxSpeed=" + strconv.Itoa(c.MaxSpeed()))
	fmt.Println("TireCount=" + strconv.Itoa(c.TireCount()))
}
