package tokenjwt

import "fmt"

func InitJWTKey(key string) {
	SecretKey = key
	fmt.Println("✅ JWT key initialized successfully")
}
