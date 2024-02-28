package main

import "fmt"

type MyString string

func (ms *MyString) ChangeName(name MyString) {

}

func main() {
	// handler.StartApp()

	var name1 MyString = "John"

	var name2 *MyString = &name1

	var name3 MyString = *name2

	fmt.Println(name2)

	fmt.Println(*name2)
}
