package sum

func Sum(slice []int) int {
	sum := 0

	for _, n := range slice {
		sum += n
	}

	return sum
}
