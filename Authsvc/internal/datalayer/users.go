package datalayer
import (
	"crypto/sha256"
	"errors"
    "log"
    
    _ "github.com/lib/pq"
)

func LoginUser(username string, password string) (User, error) {
	db := GetDB()
	if db == nil {
		return User{}, errors.New("Unable to connect to DB")
	}
	
	user := User{}
	db.Get(&user, "SELECT * from users where username=$1", username)

	sha_hasher := sha256.New()
	sha_hasher.Write([]byte(password + user.salt))
	password_hash := sha_hasher.Sum(nil)

	if string(password_hash) != user.password_hash {
		log.Print("Debug: User login failed. Username=", username)
		return User{}, errors.New("Incorrect password")
	}

	return user, nil
}
