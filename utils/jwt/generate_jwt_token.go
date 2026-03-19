package jwt

import (
	"pulse/utils/config"
	"time"

	jtoken "github.com/golang-jwt/jwt/v5"
)

type GenerateJwtTokenManager struct {
	config *config.Env
}

func NewGenerateJwtTokenManager(config *config.Env) *GenerateJwtTokenManager {
	return &GenerateJwtTokenManager{
		config: config,
	}
}

func (jwtManager *GenerateJwtTokenManager) GenerateUserJWT(id, email, role string) (string, error) {
	month := (time.Hour * 24) * 30

	claims := jtoken.MapClaims{
		"sub":   id,
		"email": email,
		"role":  role,
		"exp":   time.Now().Add(month * 6).Unix(),
	}

	token := jtoken.NewWithClaims(jtoken.SigningMethodHS256, claims)
	return token.SignedString([]byte(jwtManager.config.JwtSecretKey))
}
