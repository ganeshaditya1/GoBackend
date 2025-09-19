package datalayer

type User struct {
    userid int `db:"userid"`
    username  string `db:"last_name"`
    email_address  string `db:"email_address"`
	password_hash string `db:"password_hash"`
	salt string `db:"salt"`
	age int `db:"age"`
}