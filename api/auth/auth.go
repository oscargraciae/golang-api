package auth

import (
	"api-app/api/models"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type JWTToken struct {
	AccessToken string    `json:"accessToken"`
	ExpiresAt   time.Time `json:"expiresAt"`
}

func GenToken(user models.User) (*JWTToken, error) {
	jwtToken := jwt.New(jwt.GetSigningMethod("HS256"))
	expiresAt := time.Now().Add(time.Hour * 24 * 7) // 1 week
	jwtToken.Claims = jwt.MapClaims{
		"id":  user.ID,
		"first_name": user.FirstName,
		"last_name": user.LastName,
		"email": user.Email,
		"exp": expiresAt.Unix(),
	}

	accessToken, err := jwtToken.SignedString([]byte("JWT_SECRET"))
	if err != nil {
		return nil, err
	}

	return &JWTToken{
		AccessToken: accessToken,
		ExpiresAt:   expiresAt,
	}, nil
}