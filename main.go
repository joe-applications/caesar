package main

import (
	"fmt"
	"os"
	"github.com/joe-applications/caesar/encrypt"
	"github.com/joe-applications/caesar/decrypt"
)

func main(){
	inputStr := os.Args[1]
	encryptedStr := encrypt.Encrypt(inputStr)
	fmt.Printf("Input: %s encrypted to %s\n", inputStr, encryptedStr)
	fmt.Println("----")
	decryptedStr := decrypt.Decrypt(encryptedStr)
	fmt.Printf("Encrypted String: %s decrypted to %s\n", encryptedStr, decryptedStr)
}
