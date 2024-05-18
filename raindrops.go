package main

import (
	"fmt"
	"strconv"
)

func Convert(number int) string {
	if number%3 == 0 && number%5 == 0 && number%7 == 0 {
		return "PlingPlangPlong"
	} else if number%3 == 0 && number%5 == 0 {
		return "PlingPlang"
	} else if number%5 == 0 && number%7 == 0 {
		return "PlangPlong"
	} else if number%3 == 0 && number%7 == 0 {
		return "PlingPlong"
	} else if number%3 == 0 {
		return "Pling"
	} else if number%5 == 0 {
		return "Plang"
	} else if number%7 == 0 {
		return "Plong"
	}
	return fmt.Sprint(number)
}

func Better(number int) string {
	fizzbuzz := ""

	if number%3 == 0 {
		fizzbuzz += "Pling"
	}
	if number%5 == 0 {
		fizzbuzz += "Plang"
	}
	if number%7 == 0 {
		fizzbuzz += "Plong"
	}

	if len(fizzbuzz) == 0 {
		return strconv.Itoa(number)
	}
	return fizzbuzz

}

func main() {
	fmt.Println(Better(21))
}
