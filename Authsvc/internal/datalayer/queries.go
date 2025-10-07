package datalayer

import (
	"errors"
	"log/slog"

    _ "github.com/lib/pq"
)



func LoginUser(username string, password string) (User, error) {
	db := GetDB()
	if db == nil {
		return User{}, errors.New("Unable to connect to DB")
	}
	
	// Read the user from the DB.
	user := User{}
	err := db.Get(&user, "SELECT * FROM users WHERE username=$1", username)
	if err != nil {
		slog.Debug("Problem with reading from db.", "Error", err)
		return User{}, errors.New("Problem with reading from DB")
	}

	// Use their userid to see if they are an admin
	var count int;
	err = db.Get(&count, "SELECT COUNT(*) as count FROM admin WHERE userid=$1", user.Userid)
	if err != nil {
		slog.Debug("Problem with reading from db", "Error", err)
		return User{}, errors.New("Problem with reading from DB")
	}
	user.IsAdmin = (count > 0)

	// Hash the password using the salt of this user and 
	// see if this hash matches what we stored for them in the db.
	password_hash := hashPassword(password, user.Salt)

	// If the hash doesn't match, return an error.
	if password_hash != user.PasswordHash {
		slog.Debug("User login failed.", "Username", username)
		return User{}, errors.New("Incorrect password")
	}

	return user, nil
}


func CreateUser(user User) error{
	db := GetDB()
	if db == nil {
		return errors.New("Unable to connect to DB")
	}
	
	tx := db.MustBegin()
	_, err := tx.Exec(`INSERT INTO users(username, email_address, password_hash, salt, age) 
				VALUES($1, $2, $3, $4, $5)`, 
				user.Username, 
				user.EmailAddress, 
				user.PasswordHash, 
				user.Salt,
				user.Age)
	tx.Commit()
	if err != nil {
		return err
	}
	return nil
}
