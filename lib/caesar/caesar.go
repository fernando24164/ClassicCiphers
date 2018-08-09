package caesar

import (
	"strings"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func EncryptCaesar(message string, rotateNumber int) string {
	formatMessage := strings.ToLower(message)
	messageArray := []byte(formatMessage)
	for pos, char := range messageArray {
		if 'a' <= char && char <= 'z' {
			numberAlphabet := ((int(char) - 97) + rotateNumber) % len(alphabet)
			messageArray[pos] = alphabet[numberAlphabet]
		}
	}
	return string(messageArray)
}

func DecryptCaesar(message string, rotateNumber int) string {
	formatMessage := strings.ToLower(message)
	messageArray := []byte(formatMessage)
	for pos, char := range messageArray {
		if 'a' <= char && char <= 'z' {
			numberInAlphabet := ((int(char) - 97) - rotateNumber) % len(alphabet)
			messageArray[pos] = alphabet[numberInAlphabet]
		}
	}
	return string(messageArray)
}

func stadisticAnalysis(message string) map[string]float32 {
	formatMessage := strings.ToLower(message)
	alphabetArray := []byte(alphabet)
	messageArray := []byte(formatMessage)
	lettersSizes := make(map[string]float32, len(alphabet))
	for _, alphaChar := range alphabetArray {
		lettersSizes[string(alphaChar)] = 0.0
		countLettersCoincidences := 0
		for _, messageChar := range messageArray {
			if string(alphaChar) == string(messageChar) {
				countLettersCoincidences++
			}
		}
		lettersSizes[string(alphaChar)] = float32(countLettersCoincidences) / float32(len(formatMessage))
	}
	return lettersSizes
}

func getMaxKeyValue(frecuencies map[string]float32) string {
	var max float32
	var maxKey string
	for key := range frecuencies {
		if max <= frecuencies[key] {
			max = frecuencies[key]
			maxKey = key
		}
	}
	return maxKey
}

func BruteForceAttack(message string) int {
	frecuenciesCipherText := stadisticAnalysis(message)
	letterHighSize := getMaxKeyValue(frecuenciesCipherText)
	alphabetArray := []byte(alphabet)
	var posELetter int
	var posLetterHighSize int
	for pos, char := range alphabetArray {
		if char == 'e' {
			posELetter = pos
		}
		if string(char) == letterHighSize {
			posLetterHighSize = pos
		}
	}
	answer := (posLetterHighSize - posELetter) % len(alphabet)
	return answer
}
