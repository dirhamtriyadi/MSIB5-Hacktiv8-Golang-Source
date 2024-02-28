package main

import (
	"fmt"
	"reflect"
)

type LoginDto struct {
	Username string
	Password string
}

func main() {
	var loginDto LoginDto = LoginDto{
		Username: "user1",
		Password: "123##",
	}
	_ = loginDto

	// var num uint = 300

	var reflectType reflect.Type = reflect.TypeOf(loginDto)

	// fmt.Println(reflectType.Kind())

	// fmt.Println(reflectType.Kind() == reflect.Struct)

	// fmt.Println(reflectType.Kind() == reflect.Bool)

	// fmt.Println(reflectType.NumField())

	// fmt.Printf("%+v\n", reflectType.Field(0))

	for i := 0; i < reflectType.NumField(); i++ {
		valueOf := reflect.ValueOf(loginDto).Field(i).Interface()
		// fmt.Println(reflect.ValueOf(loginDto).Field(i))
		fmt.Printf("fieldName: %s, fieldValue: %v\n", reflectType.Field(i).Name, valueOf)

	}
}
