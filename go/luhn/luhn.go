package luhn

import (
	"regexp"
	"strconv"
)

// Valid checks if a string of numbers follows luhn's algorithm
func Valid(nums string) bool {

	//trim whitespace
	reg, err := regexp.Compile("[\\s]+")
	checkErr(err)
	processed := reg.ReplaceAllString(nums, "")

	if len(processed) < 2 {
		return false
	}

	// check that string is only numbers
	reg, err = regexp.Compile("[^0-9]+")
	checkErr(err)
	if reg.MatchString(processed) == true {
		return false
	}

	// create a flag to double every other number
	double := false

	// double every other number
	luhn := make([]int, len(processed))
	for i := len(processed) - 1; i >= 0; i-- {
		num, _ := strconv.Atoi(string(processed[i]))

		if double {
			num = doubleNum(num)
		}
		double = !double
		luhn[i] = num
	}

	// add up all the numbers
	var sum int
	for _, num := range luhn {
		sum += num
	}

	return sum%10 == 0
}

func doubleNum(num int) int {
	num *= 2
	if num > 9 {
		num = num - 9
	}
	return num
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
