package prob1

func Prob1(number int) int {
	if number-1 == 0 {
		return 0
	}

	if (number-1)%3 == 0 || (number-1)%5 == 0 {
		return (number - 1) + Prob1(number-1)
	}

	return Prob1(number - 1)
}
