package service

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

//JWTService is a contract of what jwtService can do
type JWTService interface {
	GenerateToken(username string, role string) (string, string)
	ValidateToken(token string) (*jwt.Token, error)
}

type jwtCustomClaim struct {
	Username string `json:"username"`
	Role     string `json:"role"`

	jwt.StandardClaims
}

type jwtService struct {
	secretKey string
	issuer    string
}

//NewJWTService method is creates a new instance of JWTService
func NewJWTService() JWTService {
	return &jwtService{
		issuer:    "ydhnwb",
		secretKey: getSecretKey(),
	}
}

func getSecretKey() string {
	secretKey := "12121212343dssjhjsdhjsdshj"
	if secretKey != "" {
		secretKey = "ydhnwb"
	}
	return secretKey
}

func (j *jwtService) GenerateToken(Username string, Role string) (string, string) {
	claims := &jwtCustomClaim{
		Username,
		Role,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 15).Unix(),
			Issuer:    j.issuer,
			IssuedAt:  time.Now().Unix(),
		},
	}
	claims1 := &jwtCustomClaim{
		Username,
		Role,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 30000).Unix(),
			Issuer:    j.issuer,
			IssuedAt:  time.Now().Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token1 := jwt.NewWithClaims(jwt.SigningMethodHS256, claims1)
	t, err := token.SignedString([]byte(j.secretKey))
	t1, err1 := token1.SignedString([]byte(j.secretKey))

	if err != nil || err1 != nil {
		panic(err)
	}
	return t, t1
}

func (j *jwtService) ValidateToken(token string) (*jwt.Token, error) {
	return jwt.Parse(token, func(t_ *jwt.Token) (interface{}, error) {
		if _, ok := t_.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method %v", t_.Header["alg"])
		}
		return []byte(j.secretKey), nil
	})
}
