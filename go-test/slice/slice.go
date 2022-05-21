package slice

func Slice(numbers []int) (sum int) {
	for _, number := range numbers {
		number += sum
	}
	return
}
