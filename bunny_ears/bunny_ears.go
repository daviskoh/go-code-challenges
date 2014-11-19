package bunny_ears

func BunnyEars(bunnies int) int {
	if bunnies == 0 {
		return bunnies
	}

	return 2 + BunnyEars(bunnies-1)
}
