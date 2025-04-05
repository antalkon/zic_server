package hash

import (
	"errors"
	"log"

	"github.com/matthewhartstonge/argon2"
)

func GenerateHash(password string) (string, error) {
	config := argon2.DefaultConfig()
	hash, err := config.HashEncoded([]byte(password))
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func ComparePassword(password, hash string) error {
	log.Println("🔍 Comparing Passwords:")
	log.Println("Entered Password:", password)
	log.Println("Stored Hash:", hash)

	status, err := argon2.VerifyEncoded([]byte(password), []byte(hash))
	if err != nil {
		log.Println("❌ Passwords verification failed:", err)
		return errors.New("password verification failed")
	}

	if !status {
		log.Println("❌ Passwords do not match")
		return errors.New("passwords do not match")
	}

	log.Println("✅ Passwords match")
	return nil
}
