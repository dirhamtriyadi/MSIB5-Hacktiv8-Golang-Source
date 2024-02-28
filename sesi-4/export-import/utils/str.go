package utils

import (
	"fmt"
	"sesi-4-v2/entity"
)

var RandomCharacter = "Hello World"

var secretKey = "ini rahasia"

func GetRandomNumber() {
	fmt.Println(RandomNumber)
}

func FetchProduct() entity.Product {
	return entity.Product{
		Name:  "Handphone",
		Price: 2.3,
	}
}
