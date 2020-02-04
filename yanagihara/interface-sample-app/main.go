package main

import (
	"fmt"

	"./car"
)

func main() {
	fmt.Println("test")
	normal := car.NewNormalCar()
	motorcycle := car.NewMotorcycle()
	bicycle := car.NewBicycle()
	fmt.Println("=====NormalCar")
	car.PrintCar(normal)
	fmt.Println("=====Motorcycle")
	car.PrintCar(motorcycle)
	fmt.Println("=====Bicycle")
	car.PrintCar(bicycle)
}
