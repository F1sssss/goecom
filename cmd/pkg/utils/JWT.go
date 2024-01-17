package utils

import (
	"fmt"

	"github.com/F1sssss/goecom/cmd/pkg/models"
	"github.com/golang-jwt/jwt/v5"
)

type jwtCustomClaims struct {
	Username string      `json:"username"`
	UserID   uint64      `json:"user_id"`
	Email    string      `json:"email"`
	Role     models.Role `json:"role"`
	jwt.RegisteredClaims
}

func GenerateJWT(username string, userID uint64, email string, role models.Role) (string, error) {
	fmt.Println("Generating JWT for user:", username)
	claims := &jwtCustomClaims{
		username,
		userID,
		email,
		role,
		jwt.RegisteredClaims{},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte("definetlysomethingsecrtet"))
}

func ParseJWT(tokenString string) (*jwt.Token, *jwtCustomClaims, error) {
	claims := &jwtCustomClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte("definetlysomethingsecrtet"), nil
	})
	return token, claims, err
}

func VerifyJWT(tokenString string) (bool, error) {
	token, _, err := ParseJWT(tokenString)
	if err != nil {
		return false, err
	}
	return token.Valid, nil
}

func GetJWTClaims(tokenString string) (*jwtCustomClaims, error) {

	_, claims, err := ParseJWT(tokenString)
	if err != nil {
		return nil, err
	}
	return claims, nil
}
