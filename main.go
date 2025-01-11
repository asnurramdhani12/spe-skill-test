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

func (s SpeSkillTest) ParityOutlier(number []int) string {
	var (
		// For Counter
		evenCount = 0
		oddCount  = 0

		// For Store
		lastEven = 0
		lastOdd  = 0
	)

	for _, num := range number {
		if num%2 == 0 {
			evenCount++
			lastEven = num
		} else {
			oddCount++
			lastOdd = num
		}
	}

	if evenCount > oddCount {
		if oddCount == 0 {
			return "No Outlier, All Even"
		}
		return fmt.Sprintf("Outlier: %d", lastOdd)
	}
	if evenCount < oddCount {
		if evenCount == 0 {
			return "No Outlier, All Odd"
		}
		return fmt.Sprintf("Outlier: %d", lastEven)
	}

	return "No Outlier"
}

func main() {
	// Create an instance of the SpeSkillTest struct
	speSkillTest := SpeSkillTest{}

	// Call the Narcissistic method
	fmt.Println("Output Narcissistic : ", speSkillTest.Narcissistic(1634))

	// Call the ParityOutlier method
	fmt.Println("Output ParityOutlier : ", speSkillTest.ParityOutlier([]int{2, 4, 0, 100, 4, 11, 2602, 36}))
}
