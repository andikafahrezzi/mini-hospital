package service

import (
	"backend/internal/repository"
	"backend/pkg/database"
	"errors"
	"golang.org/x/crypto/bcrypt"
	"github.com/golang-jwt/jwt/v5"
	"backend/internal/middleware"
	"time"
)


func Login(username, password string) (string, error) {

	db := database.DB

	user, roleName, err := repository.GetUserByUsername(db, username)
	if err != nil {
		return "", errors.New("invalid credentials")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	if err != nil {
		return "", errors.New("invalid credentials")
	}

	claims := jwt.MapClaims{
		"user_id":   user.ID,
		"username":  user.Username,
		"role":      roleName,
		"dokter_id": user.DokterID,
		"exp":       time.Now().Add(24 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString(middleware.GetJwtSecret())
	if err != nil {
		return "", err
	}

	return signedToken, nil
}