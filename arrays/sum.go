package arrays

func Sum(numbers []int) (result int) {
	for _, number := range numbers {
		result += number
	}
	return
}

func SumAll(numbers ...[]int) (result []int) {
	for _, number := range numbers {
		result = append(result, Sum(number))
	}
	return
}

func SumAllTails(numbers ...[]int) (result []int) {
	for _, number := range numbers {
		if len(number) == 0 {
			result = append(result, 0)
		} else {
			tailNumbers := number[1:]
			result = append(result, Sum(tailNumbers))
		}
	}
	return
}
