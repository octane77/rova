package services

import (
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"github.com/joho/godotenv"
	"os"
	"time"
)

type JwtService interface {
	GenerateToken(userId uint) string
	ValidateToken(token string) (*jwt.Token, error)
}

type jwtCustomClaims struct {
	UserId uint `json:"userId"`
	jwt.RegisteredClaims
}

type jwtService struct {
	secret string
	issuer string
}

func (j jwtService) GenerateToken(userId uint) string {
	claims := &jwtCustomClaims{
		UserId: userId,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: &jwt.NumericDate{
				Time: time.Now().Add(24 * time.Hour),
			},
			Issuer: j.issuer,
			IssuedAt: &jwt.NumericDate{
				Time: time.Now(),
			},
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(j.secret))
	if err != nil {
		panic(fmt.Sprintf("Error Creating Token: %v", err))
	}
	return t
}

func (j jwtService) ValidateToken(token string) (*jwt.Token, error) {
	return jwt.Parse(token, func(_t *jwt.Token) (interface{}, error) {
		if _, ok := _t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected Signing Method %s", _t.Header["alg"])
		}
		return []byte(j.secret), nil
	})
}

func NewJWTService() JwtService {
	return &jwtService{
		secret: getJWTSecret(),
		issuer: "ETE",
	}
}

func getJWTSecret() string {
	err := godotenv.Load()
	if err != nil {
		return ""
	}
	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		return "secret"
	}
	return jwtSecret
}
