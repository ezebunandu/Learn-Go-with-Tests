package sum

func Sum(numbers []int) (sum int) {
	if len(numbers) == 0 {
		return 0
	}
	sum = 0
	for _, num := range numbers {
		sum += num
	}
	return sum
}

func SumAll(numbersToAdd ...[]int) []int {
	result := []int{}
	for _, numbers := range numbersToAdd {
		result = append(result, Sum(numbers))
	}
	return result
}

func SumAllTails(numbersToSum ...[]int) []int {
	result := []int{}
	for _, numbers := range numbersToSum {
		if len(numbers) == 0 {
			result = append(result, 0)
		} else {
			result = append(result, Sum(numbers[1:]))
		}
	}
	return result
}
