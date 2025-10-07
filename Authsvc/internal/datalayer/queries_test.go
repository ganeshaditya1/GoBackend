package datalayer

import (
    "database/sql"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
    "github.com/jmoiron/sqlx"
)

func TestShouldCreateUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	
	SetDB(sqlx.NewDb(db, "Mock DB"))

	mock.ExpectBegin()
	mock.ExpectExec("INSERT INTO users").
		WithArgs("John", "JohnDoe@gmail.com", sqlmock.AnyArg(), sqlmock.AnyArg(), 45).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	user := NewUser("John", "JohnDoe@gmail.com", "Password", 45)
	CreateUser(user)
}

func TestShouldLoginAdminUser(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	userid, username, email_address, password_hash, password_salt, age := 55, "John", 
	"JohnDoe@gmail.com", hashPassword("Password", "Salt"), "Salt", 45
	
	SetDB(sqlx.NewDb(db, "Mock DB"))
	rows := sqlmock.NewRows([]string{"userid", "username", "email_address", "password_hash", "salt", "age"}).
			AddRow(55, "John", "JohnDoe@gmail.com", hashPassword("Password", "Salt"), "Salt", 45)

	mock.ExpectQuery("SELECT * FROM users WHERE username=$1").
		WithArgs("John").
		WillReturnRows(rows)
	

	mock.ExpectQuery("SELECT COUNT(*) as count FROM admin WHERE userid=$1").
		WithArgs(55).
		WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(1))

	user, err := LoginUser("John", "Password")
	if err != nil {
		t.Errorf("Non nil error. %v", err)
	}

	if user.Userid != userid {
		t.Errorf("Expected userid to be %d. Got %d instead", 55, user.Userid)
	}

	if user.Username != username {
		t.Errorf("Expected username to be %s. Got %s instead", "John", user.Username)		
	}

	if user.EmailAddress != email_address {
		t.Errorf("Expected email address to be %s. Got %s instead", email_address, user.EmailAddress)
	}

	if user.PasswordHash != password_hash {
		t.Errorf("Expected password hash to be %s. Got %s instead", password_hash, user.PasswordHash)
	}

	if user.Salt != password_salt {
		t.Errorf("Expected password salt to be %s. Got %s instead", password_salt, user.Salt)
	}

	if user.Age != age {
		t.Errorf("Expected age to be %d. Got %d instead", age, user.Age)
	}

	if !user.IsAdmin {
		t.Errorf("Expected user to be Admin")
	}
}

func TestShouldLoginNonAdminUser(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	userid, username, email_address, password_hash, password_salt, age := 55, "John", 
	"JohnDoe@gmail.com", hashPassword("Password", "Salt"), "Salt", 45
	
	SetDB(sqlx.NewDb(db, "Mock DB"))
	rows := sqlmock.NewRows([]string{"userid", "username", "email_address", "password_hash", "salt", "age"}).
			AddRow(55, "John", "JohnDoe@gmail.com", hashPassword("Password", "Salt"), "Salt", 45)

	mock.ExpectQuery("SELECT * FROM users WHERE username=$1").
		WithArgs("John").
		WillReturnRows(rows)
	

	mock.ExpectQuery("SELECT COUNT(*) as count FROM admin WHERE userid=$1").
		WithArgs(55).
		WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(0))

	user, err := LoginUser("John", "Password")
	if err != nil {
		t.Errorf("Non nil error. %v", err)
	}

	if user.Userid != userid {
		t.Errorf("Expected userid to be %d. Got %d instead", 55, user.Userid)
	}

	if user.Username != username {
		t.Errorf("Expected username to be %s. Got %s instead", "John", user.Username)		
	}

	if user.EmailAddress != email_address {
		t.Errorf("Expected email address to be %s. Got %s instead", email_address, user.EmailAddress)
	}

	if user.PasswordHash != password_hash {
		t.Errorf("Expected password hash to be %s. Got %s instead", password_hash, user.PasswordHash)
	}

	if user.Salt != password_salt {
		t.Errorf("Expected password salt to be %s. Got %s instead", password_salt, user.Salt)
	}

	if user.Age != age {
		t.Errorf("Expected age to be %d. Got %d instead", age, user.Age)
	}

	if user.IsAdmin {
		t.Errorf("Expected user to be not Admin")
	}
}


func TestShouldNotLoginNonExistantUser(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	
	SetDB(sqlx.NewDb(db, "Mock DB"))

	mock.ExpectQuery("SELECT * FROM users WHERE username=$1").
		WithArgs("John").
		WillReturnError(sql.ErrNoRows)

	user, err := LoginUser("John", "Password")
	if user.Username != "" {
		t.Errorf("User should not have been allowed to login. %v", user)
	}

	if err == nil {
		t.Errorf("Login should have failed and an error message should have been returned.")
	}
}