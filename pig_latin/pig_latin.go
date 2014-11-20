package pig_latin

var Vowels = [5]string{"a", "e", "i", "o", "u"}

func indexOfVowel(word string) int {
	for i, letter := range word {
		if isVowel(string(letter)) {
			return i
		}
	}

	return -1
}

func isVowel(letter string) bool {
	for _, vowel := range Vowels {
		if vowel == letter {
			return true
		}
	}

	return false
}

func PigLatin(word string) string {
	if len(word) == 0 {
		return word
	}

	startingLetter := word[0:1]
	if isVowel(startingLetter) || startingLetter == "y" {
		return word + "way"
	}

	// find index of 1st vowel
	// move string.substring(startingLetter, index) to end of string
	index := indexOfVowel(word)
	if index != -1 {
		word = word[index:] + word[0:index]
	}

	return word + "ay"
}
