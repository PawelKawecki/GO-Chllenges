package main

import (
	"fmt"
	"os"
)

var cipherText = "ZEBRASCDFGHIJKLMNOPQTUVWXYzebrascdfghijklmnopqtuvwxy";
var coresponds = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz";

func main()  {
	fmt.Println("hello challenge #3 intermediate")

	text := os.Args[1]

	//encrypted := encrypt(text)
	decrypted := decrypt(text)

	//fmt.Println(encrypted)
	fmt.Println(decrypted)
}

func encrypt(text string) string {
	encrypted := []byte(text)

	for key, val := range text  {
		for ckey, cval := range coresponds {
			if cval == val {
				encrypted[key] = cipherText[ckey]
			}
		}
	}

	return string(encrypted)
}

func decrypt(encrypted string) string {
	decrypted := []byte(encrypted)

	for key, val := range encrypted  {
		for ckey, cval := range cipherText {
			if cval == val {
				decrypted[key] = coresponds[ckey]
			}
		}
	}

	return string(decrypted)
}