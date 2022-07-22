package auth

import (
	"errors"

	"github.com/dgrijalva/jwt-go"
)

type Service interface {
	GenerateToken(userID int) (string, error)
	ValidateToken(encodedtoken string) (*jwt.Token, error)
}

type jwtService struct {
}

var SECRET_KEY = []byte("STARTUP__s3cr3T_k3Y")

func NewService() *jwtService {
	return &jwtService{}
}

func (s *jwtService) GenerateToken(userID int) (string, error) {
	claim := jwt.MapClaims{}
	claim["id"] = userID

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	signToken, err := token.SignedString(SECRET_KEY)

	if err != nil {
		return signToken, err
	}
	return signToken, nil
}

func (s *jwtService) ValidateToken(encodedtoken string) (*jwt.Token, error) {
	token, err := jwt.Parse(encodedtoken, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("Invalid Token")
		}

		return []byte(SECRET_KEY), nil

	})

	if err != nil {
		return token, err
	}

	return token, nil
}
