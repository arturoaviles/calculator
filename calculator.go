package calculator

// Calculate sums all the numbers passed to it.
func Calculate(numbers []int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}
