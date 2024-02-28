package main

import "fmt"

func main() {
	var person map[string]string = map[string]string{
		"person1": "John Doe",
		"person2": "Yolan",
		"person3": "Beny",
	}

	_ = person

	var scores map[string]int = map[string]int{
		"John Doe": 200,
		"Annabel":  140,
		"Anton":    30,
	}

	for key, value := range scores {
		fmt.Printf("key: %s, value: %d\n", key, value)
	}

	// fmt.Println("person1:", person["person1"])

	// fmt.Println("person2:", person["person2"])

	// fmt.Println("person3:", person["person3"])
}
