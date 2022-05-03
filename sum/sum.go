package sum

func Sum(numbers []int) (sum int) {
	sum = 0
	for _, num := range numbers {
		sum += num
	}
	return sum
}
