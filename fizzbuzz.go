package fizzbuzz

import "strconv"

func fizzbuzz(input int) string {
	returnString := ""
	if isDivisibleBy3(input) {
		returnString += "Fizz"
	}
	if isDivisibleBy5(input) {
		returnString += "Buzz"
	}
	if len(returnString) == 0 {
		returnString = strconv.Itoa(input)
	}
	return returnString
}

func isDivisibleBy3(input int) bool {
	return input%3 == 0
}

func isDivisibleBy5(input int) bool {
	return input%5 == 0
}
