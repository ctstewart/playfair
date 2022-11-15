package main

import (
	"fmt"

	"github.com/ctstewart/playfair/playfair"
	"github.com/pkg/errors"
)

func Example() error {
	message := "lordgranvillesletter"
	key := "palmerston"
	encrypted, err := playfair.Encrypt(message, key)
	if err != nil {
		return errors.Wrap(err, "Encrypting example")
	}
	decrypted, err := playfair.Decrypt(encrypted, key)
	if err != nil {
		return errors.Wrap(err, "Decrypting example")
	}

	fmt.Println()
	fmt.Println("--- Example ---")
	fmt.Println()
	fmt.Printf("%12s %s\n", "Message:", message)
	fmt.Printf("%12s %s\n", "Key:", key)
	fmt.Printf("%12s %s\n", "Encrypted:", encrypted)
	fmt.Printf("%12s %s\n", "Decrypted:", decrypted)
	return nil
}
