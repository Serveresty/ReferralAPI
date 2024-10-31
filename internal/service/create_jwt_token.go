package service

import (
	"ReferralAPI/models"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func (rc *RefService) CreateJWTToken(id string) (string, error) {
	claims := &models.Claims{
		StandardClaims: jwt.StandardClaims{
			Subject:   id,
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))
	if err != nil {
		return "", err
	}

	return "Bearer " + tokenString, nil
}
