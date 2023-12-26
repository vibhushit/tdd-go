package main

func Sum(numbers []int) int {
	sum := 0

	for _, number := range numbers {
		sum += number
	}
	return sum
}

func SumAll(numbersToSum ...[]int) []int {
	var res []int

	for _, numbers := range numbersToSum {
		sum := Sum(numbers)
		res = append(res, sum)
	}

	return res
}

func SumAllTails(numbersToSum ...[]int) []int {

	var res []int
	for _, numbers := range numbersToSum {
		if len(numbers) == 0 {
			res = append(res, 0)
		} else {
			newArray := numbers[1:]

			sum := Sum(newArray)

			res = append(res, sum)
		}

	}
	return res
}
