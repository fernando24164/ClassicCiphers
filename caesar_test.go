package main

import "testing"

func TestEncrypt(t *testing.T) {
	message := "Test message to check the code"
	encryptMessage := encryptCaesar(message, 5)
	if len(message) != len(encryptMessage) {
		t.Log("Incorrect encryption!")
		t.Fail()
	}
}

func TestDecrypt(t *testing.T) {
	message := "yjxy rjxxflj yt hmjhp ymj htij"
	originalMessage := "test message to check the code"
	decryptMessage := decryptCaesar(message, 5)
	if decryptMessage != originalMessage {
		t.Fail()
	}
}

func TestBruteForceAttack(t *testing.T) {
	originalMessage := "test message to check the code"
	rotateNumber := 15
	cryptMessage := encryptCaesar(originalMessage, rotateNumber)
	hackedNumber := bruteForceAttack(cryptMessage)
	if rotateNumber != hackedNumber {
		t.Fail()
	}
}
