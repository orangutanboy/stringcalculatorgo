package stringcalculator

import "strconv"
import "strings"
import "errors"

func Add(numbers string) (result int, err error) {
	if len(numbers) == 0 {
		return 0, nil
	}

	delimiter := ","

	if strings.HasPrefix(numbers, "//") {
		numbers = strings.TrimPrefix(numbers, "//")
		delimiter = string(numbers[0])
		numbers = strings.TrimPrefix(numbers, delimiter)
	}

	numbers = strings.Replace(numbers, "\n", delimiter, -1)
	allNumberStrings := strings.Split(numbers, delimiter)

	valid, invalidNumbers := tryValidateNumbers(allNumberStrings)
	if !valid {
		err = errors.New("negatives not allowed " + invalidNumbers)
		return
	}

	allnumbers := removeTooLarge(allNumberStrings)
	return sumNumbers(allnumbers), nil
}

func removeTooLarge(allnumbers []string) (validNumbers []int) {
	validNumbers = []int{}
	for _, number := range allnumbers {
		if current, _ := strconv.Atoi(number); current <= 1000 {
			validNumbers = append(validNumbers, current)
		}
	}
	return
}

func tryValidateNumbers(allnumbers []string) (result bool, invalidNumbers string) {
	invalidNumbers = ""
	result = true
	for _, current := range allnumbers {
		if strings.HasPrefix(current, "-") {
			invalidNumbers += current
			result = false
		}
	}
	return
}

func sumNumbers(allnumbers []int) int {
	total := 0
	for _, current := range allnumbers {
		total += current
	}
	return total
}
