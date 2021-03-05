package main

import (
	"crypto/hmac"
	"crypto/sha512"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"log"
)

var key []byte

func main() {
	for i := 1; i <= 64; i++ {
		key = append(key, byte(i))
	}
	pass := "123456789"
	hashedPass, err := hashPassword(pass)
	if err != nil {
		log.Fatalf("got an error while hashing password: %v\n", err)
	}

	err = comparePassword(pass, hashedPass)
	if err != nil {
		log.Fatalln("Not logged in  ")
	}
	log.Println("Logged In!")

}

func hashPassword(password string) ([]byte, error) {
	bs, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("error while generating bcrypt hash from password: %w", err)
	}
	return bs, nil
}

func comparePassword(password string, hashedPass []byte) error {
	err := bcrypt.CompareHashAndPassword(hashedPass, []byte(password))
	if err != nil {
		return fmt.Errorf("invalid password: %w", err)
	}
	return nil
}

func signMessage(msg []byte) ([]byte, error) {
	h := hmac.New(sha512.New, key)

	_, err :=  h.Write(msg)
	if err != nil {
		return nil, fmt.Errorf("error in singMessage while hashing message: %w\n", err)
	}

	signature := h.Sum(nil)
	return signature, nil
}

func checkSig(msg , sig []byte) (bool,error) {
	newSig, err  := signMessage(msg)
	if err != nil {
		return false, fmt.Errorf("error in checkSig while getting signature of message: %w\n", err)
	}

	same := hmac.Equal(newSig, sig)
	return same, nil
}
