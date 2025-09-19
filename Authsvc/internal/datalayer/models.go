package datalayer

type User struct {
    Userid int `db:"userid"`
    Username  string `db:"username"`
    EmailAddress  string `db:"email_address"`
	PasswordHash string `db:"password_hash"`
	Salt string `db:"salt"`
	Age int `db:"age"`
	IsAdmin bool
}