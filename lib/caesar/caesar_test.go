package caesar

import (
	caesar "ClassicCiphers/lib/caesar"
	"testing"
)

func TestEncrypt(t *testing.T) {
	message := "Test message to check the code"
	encryptMessage := caesar.EncryptCaesar(message, 5)
	if len(message) != len(encryptMessage) {
		t.Log("Incorrect encryption!")
		t.Fail()
	}
}

func TestDecrypt(t *testing.T) {
	message := "yjxy rjxxflj yt hmjhp ymj htij"
	originalMessage := "test message to check the code"
	decryptMessage := caesar.DecryptCaesar(message, 5)
	if decryptMessage != originalMessage {
		t.Fail()
	}
}

func TestBruteForceAttack(t *testing.T) {
	originalMessage := "test message to check the code"
	rotateNumber := 15
	cryptMessage := caesar.EncryptCaesar(originalMessage, rotateNumber)
	hackedNumber := caesar.BruteForceAttack(cryptMessage)
	if rotateNumber != hackedNumber {
		t.Fail()
	}
}
