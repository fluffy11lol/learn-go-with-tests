package arrays_and_slices

func Sum(numbers []int) int {
	sum := 0
	for i := range numbers {
		sum += numbers[i]
	}
	return sum
}

func SumAll(numbersToSum ...[]int) []int {
	sums := make([]int, len(numbersToSum))
	for i := range numbersToSum {
		for j := range numbersToSum[i] {
			sums[i] += numbersToSum[i][j]
		}
	}
	return sums
}
func SumAllTails(numbersToSum ...[]int) []int {
	sums := make([]int, len(numbersToSum))
	for i, tail := range numbersToSum {
		if len(tail) == 0 {
			continue
		}
		sums[i] = SumAll(tail[1:])[0]
	}
	return sums
}
