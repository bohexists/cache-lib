package main

func main() {
	println(findMin(5144, 95445, 4535, 52423, 5234))

}

func findMin(numbers ...int) int {
	if len(numbers) == 0 {
		return 0
	}

	min := numbers[0]

	for _, i := range numbers {
		if i < min {
			min = i
		}
	}
	return min
}
S