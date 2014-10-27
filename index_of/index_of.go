package index_of

func IndexOf(mainString, substring string, currentIndex int) int {
	if currentIndex+len(substring) > len(mainString) {
		return -1
	}

	// NOTE: golang taking substring of string w/ [a:b]
	// will take all elements between range a & b
	// ex. "dude"[0:4] //=> "dude"
	// ex. "dude"[0:3] //=> "dud"
	if mainString[currentIndex:currentIndex+len(substring)] == substring {
		return currentIndex
	}

	return IndexOf(mainString, substring, currentIndex+1)
}
