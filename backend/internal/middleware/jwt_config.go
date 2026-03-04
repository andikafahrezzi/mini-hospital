package middleware

import "os"

func GetJwtSecret() []byte {
	return []byte(os.Getenv("JWT_SECRET"))
}