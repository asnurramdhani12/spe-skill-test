package main

import (
	"fmt"
	"math"
	"strconv"
)

// Instead of Class in OOP Language, we use Struct in Go
type SpeSkillTest struct{}

func (s SpeSkillTest) Narcissistic(number int) bool {
	// Get length of number
	length := len(strconv.Itoa(number))

	// Convert number to string
	numberString := strconv.Itoa(number)

	// Calculate the sum of the digits
	sum := 0.0
	for _, char := range numberString {
		digit, _ := strconv.Atoi(string(char))
		sum += math.Pow(float64(digit), float64(length))
	}

	// Check if the sum is equal to the number
	return sum == float64(number)
}

func main() {
	// Create an instance of the SpeSkillTest struct
	speSkillTest := SpeSkillTest{}

	// Call the Narcissistic method
	fmt.Println(speSkillTest.Narcissistic(1634))
}
