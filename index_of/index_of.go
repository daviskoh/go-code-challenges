package index_of

func IndexOf(mainString, substring string, currentIndex int) int {
	if currentIndex+len(substring) > len(mainString) {
		return -1
	}

	if mainString[currentIndex:currentIndex+len(substring)] == substring {
		return currentIndex
	}

	return IndexOf(mainString, substring, currentIndex+1)
}
