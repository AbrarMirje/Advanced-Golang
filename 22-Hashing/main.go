package main

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"io"
)

func main() {
	password := "password123"
	// hash := sha256.Sum256([]byte(password))
	// hash512 := sha512.Sum512([]byte(password))
	// fmt.Println(password)
	// fmt.Println(hash)
	// fmt.Printf("SHA-256 Has hex val: %x\n", hash)
	// fmt.Printf("SHA-512 Has hex val: %x\n", hash512)

	salt, err := generateSalt()
	if err != nil {
		fmt.Println("error: while generating salt")
		return
	}

	// Hash the password with salt
	signUpHash := hashPassword(password, salt)

	// Store the salt and password in database, just printing now
	saltStr := base64.StdEncoding.EncodeToString(salt)
	fmt.Println("Salt: ", saltStr)
	fmt.Println("Hash: ", signUpHash)

	// verify
	// retrieve the saltStr and decode it
	decodedSalt, err := base64.StdEncoding.DecodeString(saltStr)
	if err != nil {
		fmt.Println("error: while decoding the password")
	}
	loginHash := hashPassword(password, decodedSalt)

	// compare the stored signUpHash with loginHash
	if signUpHash == loginHash {
		fmt.Println("Login successful.")
	} else {
		fmt.Println("Login Failed.")
	}

}

func generateSalt() ([]byte, error) {
	salt := make([]byte, 16)
	_, err := io.ReadFull(rand.Reader, salt)
	if err != nil {
		return nil, err
	}
	return salt, nil
}

// function to hash password
func hashPassword(password string, salt []byte) string {
	saltedPassword := append(salt, []byte(password)...)
	hash := sha256.Sum256(saltedPassword)
	return base64.StdEncoding.EncodeToString(hash[:])
}
