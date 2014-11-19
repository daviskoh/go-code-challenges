package bunny_ears_2

func BunnyEars2(bunnies int) int {
	if bunnies == 0 {
		return 0
	}

	if bunnies%2 == 0 {
		return 3 + BunnyEars2(bunnies-1)
	}

	return 2 + BunnyEars2(bunnies-1)
}
