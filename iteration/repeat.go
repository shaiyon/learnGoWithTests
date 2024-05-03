package iteration

const defaultRepeatCount = 5

func Repeat(c string, r int) (repeated string) {
	if r == 0 {
		r = defaultRepeatCount
	}
	for i := 0; i < r; i++ {
		repeated += c
	}
	return
}