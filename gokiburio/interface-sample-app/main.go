package main

import (
	"fmt"

	"github.com/uigobrik/homeworks/gokiburio/interface-sample-app/animal"
)

func main() {
	fmt.Println("test")
	// cat := &animal.CatValue{
	// 	Value: AnimalStatus{
	// 		Name: "ポポ",
	// 		Age:  2,
	// 		Type: "マンチカン",
	// 	},
	// }
	cat := &animal.CatValue{}
	cat.Value.Name = "ポポ"
	cat.Value.Age = 2
	cat.Value.Type = "マンチカン"
	animal.GetAnimalStatus(cat)
}
