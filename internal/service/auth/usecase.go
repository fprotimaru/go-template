package auth

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

type UseCase interface {
	HashPassword(passwd []byte) string
	ComparePasswords(hashedPasswd, passwd2 []byte) bool
	GenerateToken(claims jwt.Claims) (string, error)
	Middleware(next http.HandlerFunc) fiber.Handler
}
