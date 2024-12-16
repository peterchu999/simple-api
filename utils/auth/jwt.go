package auth

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

const ISSUER_APPLICATION = "simple-api"
const LOGIN_EXPIRATION_DURATION = time.Duration(5) * time.Minute

var JWT_SIGNING_METHOD = jwt.SigningMethodHS256
var JWT_SIGNATURE_KEY = []byte(os.Getenv("JWT_SIGNATURE_KEY"))

type Claim struct {
	jwt.Claims
	Username string
}

func GenerateToken(username string) (string, error) {
	claim := Claim{
		Claims: jwt.RegisteredClaims{
			Issuer:    ISSUER_APPLICATION,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(LOGIN_EXPIRATION_DURATION)),
		},
		Username: username,
	}
	unsignedToken := jwt.NewWithClaims(JWT_SIGNING_METHOD, claim)
	return unsignedToken.SignedString(JWT_SIGNATURE_KEY)
}

func ValidateToken(c *gin.Context) error {
	authorizationHeader := c.Request.Header.Get("Authorization")
	if !strings.Contains(authorizationHeader, "Bearer") {
		return errors.New("Invalid Token")
	}
	tokenString := strings.Replace(authorizationHeader, "Bearer ", "", -1)
	fmt.Println(tokenString)
	token, errParse := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if method, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Signing method invalid")
		} else if method != JWT_SIGNING_METHOD {
			return nil, fmt.Errorf("Signing method invalid")
		}
		return JWT_SIGNATURE_KEY, nil
	})

	if errParse != nil {
		return errParse
	}

	claims, errClaim := token.Claims.(jwt.MapClaims)

	if !errClaim || !token.Valid {
		return errors.New("Invalid Token")
	}
	log.Println("This is token: ", claims)

	return nil
}
