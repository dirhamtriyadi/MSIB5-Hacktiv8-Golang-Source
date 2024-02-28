package main

import (
	"errors"
	"fmt"
)

type AcMilanCallback func(string, uint) bool

func main() {

	// fmt.Println(greet("Hello everyone!!!", 12))

	// value1, value2 := greet("Hello everyone!!!", 12)

	// fmt.Println("value1:", value1)

	// fmt.Println("value2:", value2)

	// var players []string = []string{
	// 	"Nesta", "Maldini", "Ronaldo",
	// }

	// printFootballPlayers("ac milan", 100000, players...)

	// message, err := isValidAcMilanPlayer("Crespo")

	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println(message)
	// }

	// acMilanCallBack(5, isValidName)

	var callback AcMilanCallback = func(name string, nameLength uint) bool {
		nameToRune := []rune(name)
		if len(nameToRune) != int(nameLength) {
			return true
		}

		return false
	}

	callback("", 10)

	_ = callback

	acMilanCallBack(6, func(name string, nameLength uint) bool {
		nameToRune := []rune(name)
		if len(nameToRune) != int(nameLength) {
			return true
		}

		return false
	})

	acMilanCallBack(6, func(name string, nameLength uint) bool {
		nameToRune := []rune(name)
		if len(nameToRune) == int(nameLength) {
			return true
		}

		return false
	})

}

// ... => ellipsis

func isValidName(name string, nameLength uint) bool {

	nameToRune := []rune(name)
	if len(nameToRune) == int(nameLength) {
		return true
	}

	return false
}

func acMilanCallBack(nameLength uint, cb AcMilanCallback) {
	var acMilanPlayers []string = []string{"Crespo", "Nesta", "Maldini"}

	for _, eachPlayer := range acMilanPlayers {
		if cb(eachPlayer, nameLength) {
			fmt.Printf("%s has a valid name\n", eachPlayer)
		} else {
			fmt.Printf("%s has an invalid name\n", eachPlayer)
		}
	}
}

func printFootballPlayers(team string, totalBalance int, footballPlayers ...string) {
	fmt.Println("team:", team)
	fmt.Printf("footballPlayers: %#v\n", footballPlayers)
}

func isValidAcMilanPlayer(player string) (string, error) {
	var acMilanPlayers []string = []string{"Crespo", "Nesta", "Maldini"}

	var flag bool = false

	for _, eachPlayer := range acMilanPlayers {
		if player == eachPlayer {
			flag = true
			break
		}
	}

	if flag == true {
		return fmt.Sprintf("%s is a valid AC Milan player", player), nil
	}

	return "", errors.New("invalid")
}

func greet(message string, favoriteNumber int) (string, map[string]string) {
	// fmt.Println("message:", message)

	var resultString string = fmt.Sprintf("message: %s, favoriteNumber: %d", message, favoriteNumber)

	var resultMap map[string]string = map[string]string{
		"param1": message,
		"param2": fmt.Sprintf("%d", favoriteNumber),
	}

	return resultString, resultMap
}
