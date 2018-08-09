package main

import (
	caesar "ClassicCiphers/lib/caesar"
	"fmt"
)

func main() {
	message := "Secret message"
	encryptMessage := caesar.EncryptCaesar(message, 5)
	fmt.Println(encryptMessage)
	rotateNumber := caesar.BruteForceAttack(encryptMessage)
	fmt.Println(rotateNumber)
}
