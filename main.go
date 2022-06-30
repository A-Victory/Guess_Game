package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	comNum := genRandNum()
	//fmt.Println("Random number generated is: ", comNum)

	/*input := userInput()
	fmt.Println(input)*/

	for matching := false; !matching; {

		guess := userInput()

		matching := isMatching(comNum, guess)
		fmt.Println(matching)
	}
}

func genRandNum() int {
	rand.Seed(time.Now().UnixMicro())
	return rand.Int() % 11
}

func userInput() int {
	fmt.Println("Guess should be within 1 - 10")
	fmt.Println("PLease input ur guess: ")
	var input int
	_, err := fmt.Scan("&input")
	if err != nil {
		fmt.Println("Failed to Parse")
	} else {
		fmt.Println("YOU GUESSED: ", input)
	}
	return input
}

func isMatching(comNum int, input int) bool {
	if comNum > input {
		fmt.Println("Your guess is too high")
		return false
	} else if comNum < input {
		fmt.Println("Your guess is too low")
		return false
	} else {
		fmt.Println("YOU GUESSED RIGHT!")
		return true
	}
}
