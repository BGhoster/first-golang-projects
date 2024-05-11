package main

import (
	"fmt"
)

// Adding 2 numbers together
func adding_numbers(number_1 float32, number_2 float32) float32 {
	return number_1 + number_2
}

// Subtracting 2 numbers
func subtracting_numbers(n1 float32, n2 float32) float32 {
	return n1 - n2
}

func dividing_numbers(n1 float32, n2 float32) float32 {
	return n1 / n2
}

func main() {
	fmt.Println(adding_numbers(4, 2))
	fmt.Printf("%s %f \n", "Subtraction", subtracting_numbers(5, 2))
	fmt.Println(dividing_numbers(60, 6))
}
