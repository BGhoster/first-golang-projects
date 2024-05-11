package main

import (
	"fmt"
	"math/rand"
)

func game() {
	fmt.Println("Welcome to the random guessing game")
	fmt.Println("You need to guess the number the computer generates 0-10")
	// Computer choice
	computer_choice := rand.Intn(10)

	// User choice
	fmt.Print("Enter a number ")
	var user_choice int
	fmt.Scan(&user_choice)

	if computer_choice == user_choice {
		fmt.Println("You got it")
	} else {
		fmt.Printf("%s %d", "Nope", computer_choice)
	}
}

func main() {
	game()
}
