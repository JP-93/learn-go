package interation

const repeatCount = 5

func Repeat(char string) (repeated string) {
	for i := 0; i < repeatCount; i++ {
		repeated += char
	}
	return
}
