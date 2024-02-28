package main

import "fmt"

func main() {
	var greet string = "Hello"

	// var car string = "BMW"

	// const carType string = car

	const car string = "BMW"

	var carType string = car

	_ = carType

	// for i := 0; i < len(greet); i++ {
	// 	var char string = string(greet[i])

	// 	fmt.Println(char)
	// }

	var city string = "oraș"

	var cityRune []rune = []rune(city)

	_ = cityRune

	fmt.Println("byte =>", len(city))
	fmt.Println("rune =>", len([]rune(city)))

	// for i := 0; i < len(cityRune); i++ {
	// 	fmt.Println(cityRune[i])
	// }

	// for _, eachCharacter := range city {
	// 	fmt.Println(string(eachCharacter))
	// }

	// for i := 0; i < len(city); i++ {
	// 	// var char string = string(city[i])

	// 	fmt.Println(city[i])
	// }

	// fmt.Println([]byte(city))

	var name string = "John Doe"

	var age int8 = 20

	var result string = fmt.Sprintf("%s %s, umur %d tahun", greet, name, age)

	var result2 = string(72)

	h := string(greet[0])

	e := string(greet[1])

	l := string(greet[2])

	_, _, _, _, _ = result, h, e, l, result2

	// fmt.Println(string([]byte{72, 101, 108, 108, 111}))

}
