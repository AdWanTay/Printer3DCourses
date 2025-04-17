package utils

import (
	"Printer3DCourses/internal/config"
	"Printer3DCourses/internal/models"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

func GenerateJWT(user *models.User, cfg *config.Config) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["user_id"] = user.ID
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	tokenString, _ := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(cfg.JWTSecret))
	return tokenString, nil
}
