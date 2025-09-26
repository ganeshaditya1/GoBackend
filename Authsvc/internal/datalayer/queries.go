package datalayer
import (
	"crypto/rand"
	"crypto/sha256"
    "encoding/hex"
	"errors"
	"log/slog"
    
    _ "github.com/lib/pq"

)

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

func LoginUser(username string, password string) (User, error) {
	db := GetDB()
	if db == nil {
		return User{}, errors.New("Unable to connect to DB")
	}
	
	// Read the user from the DB.
	user := User{}
	err := db.Get(&user, "SELECT * FROM users WHERE username=$1", username)
	if err != nil {
		slog.Debug("Problem with reading from db. %v", err)
		return User{}, errors.New("Problem with reading from DB")
	}

	// Use their userid to see if they are an admin
	var count int;
	err = db.Get(&count, "SELECT COUNT(*) as count FROM admin WHERE userid=$1", user.Userid)
	if err != nil {
		slog.Debug("Problem with reading from db. %v", err)
		return User{}, errors.New("Problem with reading from DB")
	}
	user.IsAdmin = (count > 0)

	// Hash the password using the salt of this user and 
	// see if this hash matches what we stored for them in the db.
	password_hash := hashPassword(password, user.Salt)

	// If the hash doesn't match, return an error.
	if password_hash != user.PasswordHash {
		slog.Debug("User login failed. Username=%s", username)
		return User{}, errors.New("Incorrect password")
	}

	return user, nil
}


func CreateUser(username string, email_address string, password string, age int) error{
	db := GetDB()
	if db == nil {
		return errors.New("Unable to connect to DB")
	}

	// Create password salt.
	password_salt := generateSalt(255)
	password_hash := hashPassword(password, password_salt)
	
	tx := db.MustBegin()
	tx.MustExec(`INSERT INTO users(username, email_address, password_hash, salt, age) 
				VALUES($1, $2, $3, $4, $5)`, 
				username, email_address, password_hash, password_salt, age)
	tx.Commit()
	return nil
}
