package util

import (
	"os"
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func TestMain(m *testing.M) {
	// ensure JWT_SECRET is set for NewJWTHandler
	os.Setenv("JWT_SECRET", "test-secret")
	code := m.Run()
	os.Unsetenv("JWT_SECRET")
	os.Exit(code)
}

func makeHandlerWithSecret(secret string) JWTHandler {
	return JWTHandler{secretKey: []byte(secret)}
}

func TestGenerateAndDecodeToken_Success(t *testing.T) {
	h := makeHandlerWithSecret("test-secret")

	claims := JWTUserClaims{
		Username: "alice",
		IsAdmin:  true,
		Email:    "a@example.com",
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "authsvc",
			Subject:   "42",
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(1 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := h.GenerateToken(claims)
	if token == "" {
		t.Fatal("expected non-empty token")
	}

	got, err := h.DecodeToken(token)
	if err != nil {
		t.Fatalf("DecodeToken error: %v", err)
	}
	if got.Username != claims.Username {
		t.Fatalf("Username mismatch: got %q want %q", got.Username, claims.Username)
	}
	if got.Email != claims.Email {
		t.Fatalf("Email mismatch: got %q want %q", got.Email, claims.Email)
	}
	if got.IsAdmin != claims.IsAdmin {
		t.Fatalf("IsAdmin mismatch: got %v want %v", got.IsAdmin, claims.IsAdmin)
	}
	if got.Subject != claims.Subject {
		t.Fatalf("Subject mismatch: got %q want %q", got.Subject, claims.Subject)
	}
}

func TestDecodeToken_InvalidSignature(t *testing.T) {
	// generate token with one secret, try decode with another
	h1 := makeHandlerWithSecret("secret-one")
	h2 := makeHandlerWithSecret("secret-two")

	claims := JWTUserClaims{
		Username: "bob",
		IsAdmin:  false,
		Email:    "b@example.com",
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "authsvc",
			Subject:   "7",
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(1 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := h1.GenerateToken(claims)
	if token == "" {
		t.Fatal("expected non-empty token")
	}
	_, err := h2.DecodeToken(token)
	if err == nil {
		t.Fatal("expected error decoding token with wrong signature, got nil")
	}
}

func TestDecodeToken_Expired(t *testing.T) {
	h := makeHandlerWithSecret("test-secret")

	claims := JWTUserClaims{
		Username: "carol",
		IsAdmin:  false,
		Email:    "c@example.com",
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "authsvc",
			Subject:   "9",
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(-1 * time.Hour)), // already expired
			IssuedAt:  jwt.NewNumericDate(time.Now().Add(-2 * time.Hour)),
		},
	}

	token := h.GenerateToken(claims)
	if token == "" {
		t.Fatal("expected non-empty token")
	}
	_, err := h.DecodeToken(token)
	if err == nil {
		t.Fatal("expected error for expired token, got nil")
	}
}

func TestNewJWTHandler_EnvVarRequirement(t *testing.T) {
	// ensure constructor reads env; temporarily set and unset
	os.Setenv("JWT_SECRET", "env-secret")
	h := NewJWTHandler()
	if string(h.secretKey) != "env-secret" {
		t.Fatalf("unexpected secret: %q", string(h.secretKey))
	}
	os.Unsetenv("JWT_SECRET")

	// restore for other tests (TestMain covers lifecycle)
	os.Setenv("JWT_SECRET", "test-secret")
}
