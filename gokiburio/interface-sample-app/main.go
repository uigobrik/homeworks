package main

import (
	"fmt"

	"github/uigobrik/homeworks/gokiburio/interface-sample-app/animal"
)

func main() {
	fmt.Println("test")
	cat := animal.Cat{Name: "ポポ",Age: 2, Type: "マンチカン"}
	animal.GetAnimalStatus(cat)
}
