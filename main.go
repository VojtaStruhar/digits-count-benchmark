package main

import (
	"strconv"
)

func main() {
	println(numericSolution(100))
}

// This function decides if a number has more even digits than odd by converting the number into a string
// and looking at each letter-number separately
// Returns true if there is more even numbers than odd
func textSolution(number int) bool {
	text := strconv.Itoa(number)

	evenScore := 0

	for _, letter := range text {
		digit, _ := strconv.Atoi(string(letter))
		if digit%2 == 0 {
			evenScore++
		} else {
			evenScore--
		}
	}

	return evenScore > 0
}

// this method takes a shortcut. Because the letter '0' is an odd number in the ascii table,
// we can look directly on the value of the byte and skip the second string conversion.
func textSolutionShortcut(number int) bool {
	text := strconv.Itoa(number)

	evenScore := 0

	for _, letter := range text {
		if letter%2 == 0 {
			evenScore++
		} else {
			evenScore--
		}
	}

	return evenScore > 0
}

// This is the correct solution (in my opinion).
func numericSolution(number int) bool {
	if number == 0 {
		return true
	}

	evenScore := 0
	for number != 0 {
		if number%2 == 0 {
			evenScore++
		} else {
			evenScore--
		}

		number /= 10
	}
	return evenScore > 0
}
