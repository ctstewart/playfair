package main

import (
	"fmt"

	"github.com/ctstewart/playfair/playfair"
	"github.com/pkg/errors"
)

func Input() (err error) {
	var eord string
	var message string
	var key string
	var ans string

	fmt.Println("Would you like to (e)ncrypt or (d)ecrypt?")
	fmt.Scanln(&eord)

	if eord != "e" && eord != "d" {
		return errors.New("please enter 'e' for encrypting or 'd' for decrypting")
	}

	fmt.Println("Please enter the key you would like to use:")
	fmt.Scanln(&key)

	if key == "" {
		return errors.New("please enter a valid key")
	}

	if eord == "e" {
		fmt.Println("Please enter the message you would like to encrypt:")
		fmt.Scanln(&message)
		ans, err = playfair.Encrypt(message, key)
		if err != nil {
			return errors.Wrap(err, "Encrypting message")
		}
	} else if eord == "d" {
		fmt.Println("Please enter the message you would like to decrypt:")
		fmt.Scanln(&message)
		ans, err = playfair.Decrypt(message, key)
		if err != nil {
			return errors.Wrap(err, "Decrypting message")
		}
	}

	fmt.Println(ans)

	return nil
}
