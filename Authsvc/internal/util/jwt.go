package util

import (
	"errors"
	"fmt"
	"os"
	"log/slog"
	"time"
	"github.com/golang-jwt/jwt/v5"
)

type JWTHandler struct {
	// secretKey should match the key used to sign tokens
	secretKey []byte
}

func NewJWTHandler() JWTHandler {
	secretkey := os.Getenv("JWT_SECRET")
	if secretkey == "" {
		slog.Error("JWT secret not specified.")
		os.Exit(1)
	}
	return JWTHandler{secretKey: []byte(secretkey)}
}

type JWTUserClaims struct {
	Username string 
	IsAdmin  bool   
	Email string 
	jwt.RegisteredClaims
}

func NewJWTUserClaims(userid int, username string, 
	emailaddress string, isadmin bool) JWTUserClaims {
	return JWTUserClaims{
		username,
		isadmin,
		emailaddress,
		jwt.RegisteredClaims {
			Issuer: "authsvc",
			Subject: string(userid),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(72 * time.Hour)), // 3 days
			IssuedAt: jwt.NewNumericDate(time.Now()),
		},
	}
}


func (jwtHandler JWTHandler) GenerateToken(userClaims JWTUserClaims) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, userClaims)

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString(jwtHandler.secretKey)
	if err != nil {
		slog.Error("Failed to produce a tokenString", err.Error())
	}
	return tokenString
}

func (jwtHandler JWTHandler) DecodeToken(token string) (JWTUserClaims, error) {
	claims := &JWTUserClaims{}
    jwtToken, err := jwt.ParseWithClaims(token, claims, func(t *jwt.Token) (any, error) {
        // Optional: check signing method
        if t.Method.Alg() != jwt.SigningMethodHS256.Alg()  {
            return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
        }
        return jwtHandler.secretKey, nil
    })
    if (err != nil || !jwtToken.Valid) {
		return JWTUserClaims{}, err
	}

	if time.Now().After(claims.ExpiresAt.Time) {
		return JWTUserClaims{}, errors.New("Token Expired.")
	}

	return *claims, nil      
}