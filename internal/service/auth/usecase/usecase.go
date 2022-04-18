package authusecase

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

type UseCase struct {
	SignedKey []byte
}

func NewUseCase(SignedKey []byte) *UseCase {
	return &UseCase{SignedKey: SignedKey}
}

func (u UseCase) HashPassword(passwd []byte) string {
	p, _ := bcrypt.GenerateFromPassword(passwd, 7)
	return string(p)
}

func (u UseCase) ComparePasswords(hashedPasswd, passwd2 []byte) bool {
	err := bcrypt.CompareHashAndPassword(hashedPasswd, passwd2)
	return err == nil
}

func (u UseCase) GenerateToken(claims jwt.Claims) (string, error) {
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
	token, err := jwtToken.SignedString(u.SignedKey)
	return token, err
}

func (u UseCase) Middleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.Next()
	}
}
