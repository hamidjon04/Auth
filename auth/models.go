package auth

import "github.com/golang-jwt/jwt/v5"

type Claim struct{
	Items map[string]interface{} `json:"items"`
	jwt.RegisteredClaims
}