package pig_latin

import (
	"strings"
)

var Vowels = "aeiou"

func indexOfVowel(word string) int {
	for i, letter := range word {
		if strings.Contains(Vowels, string(letter)) {
			return i
		}
	}

	return -1
}

func PigLatin(word string) string {
	if len(word) == 0 {
		return word
	}

	// find index of 1st vowel
	index := indexOfVowel(word)

	if index == 0 || word[0:1] == "y" {
		return word + "way"
	}

	if index != -1 {
		// move string.substring(startingLetter, index) to end of string
		word = word[index:] + word[0:index]
	}

	return word + "ay"
}
