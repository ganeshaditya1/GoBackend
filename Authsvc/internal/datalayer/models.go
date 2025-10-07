package datalayer

import (
	"crypto/rand"
	"crypto/sha256"
    "encoding/hex"
)

type User struct {
    Userid int `db:"userid"`
    Username  string `db:"username"`
    EmailAddress  string `db:"email_address"`
	PasswordHash string `db:"password_hash"`
	Salt string `db:"salt"`
	Age int `db:"age"`
	IsAdmin bool
}

func hashPassword(password string, salt string) string {
    data := []byte(password + salt)

    // Compute SHA-256 hash
    hash := sha256.Sum256(data)

    // Convert the hash to a hexadecimal string
	return hex.EncodeToString(hash[:])
}

func generateSalt(length int) string {
    // Create a byte slice to hold the random bytes
    bytes := make([]byte, length)

    // Read random bytes from the crypto/rand package
    rand.Read(bytes)

    // Convert bytes to a UTF-8 string
    // Here we will limit the byte values to printable ASCII characters
    for i := 0; i < length; i++ {
        bytes[i] = byte(32 + (bytes[i] % 95)) // Printable ASCII range from 32 to 126
    }

    return string(bytes)
}

func NewUser(username string, email string, password string, age int) User {
	// Create password salt.
	password_salt := generateSalt(255)
	password_hash := hashPassword(password, password_salt)

	return User{
		Username: username,
		EmailAddress: email,
		PasswordHash: password_hash,
		Salt: password_salt,
		Age: age,
	}
}