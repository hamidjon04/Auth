package auth

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateJWT(jwtKey string, items map[string]interface{}, expiry time.Duration)(string, error){
	exp := time.Now().Add(expiry)
	claims := Claim{
		Items: items,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(exp),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(jwtKey))
}

func ValidateJWT(jwtKey string, tokenStr string)(*Claim, error){
	claims := &Claim{}
	token, err := jwt.ParseWithClaims(tokenStr, claims, func(t *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil || !token.Valid{
		return nil, err
	}
	return claims, nil
}