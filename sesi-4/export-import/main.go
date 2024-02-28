package main

import (
	"fmt"
	"sesi-4-v2/entity"
)

func main() {

	fmt.Println()
	// fmt.Println(utils.RandomCharacter)
	// fmt.Println(utils.RandomNumber)

	// utils.GetRandomNumber()

	// fmt.Println(utils.secretkey)

	var p1 entity.Product = entity.Product{
		Name:  "Headphone",
		Price: 1.5,
	}

	fmt.Println("owner name:", p1.GetOwnerName())

	// fmt.Printf("%+v\n", p1)
}
