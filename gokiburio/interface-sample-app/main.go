package main

import (
	"fmt"

	"github.com/uigobrik/homeworks/gokiburio/interface-sample-app/animal"
)

func main() {
	fmt.Println("test")
	cat := &animal.CatValue{
		AnimalStatus: {
			Name: "ポポ",
			Age: 2,
			Type: "マンチカン",
		},
	}
	animal.GetAnimalStatus(cat)
}
