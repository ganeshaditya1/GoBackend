package datalayer

import (
	"errors"
	"log/slog"

    _ "github.com/lib/pq"
)

func CreateBook(book Book) error {
	db := GetDB()
	if db == nil {
		return errors.New("Unable to connect to DB")
	}
	
	tx := db.MustBegin()
	_, err := tx.Exec(`INSERT INTO books(book_name, book_description) 
				VALUES($1, $2)`, 
				book.Name,
				book.Description)
	tx.Commit()
	if err != nil {
		return err
	}
	return nil
}

func RemoveBook(bookid int) error {
	db := GetDB()
	if db == nil {
		return errors.New("Unable to connect to DB")
	}
	
	tx := db.MustBegin()
	_, err := tx.Exec(`DELETE FROM BOOKS WHERE book_id=$1`, bookid)
	tx.Commit()
	if err != nil {
		return err
	}
	return nil
}

func SearchBooks(bookname string) ([]Book, error) {
	db := GetDB()
	if db == nil {
		return []Book{}, errors.New("Unable to connect to DB")
	}
	
	books := []Book{}
	err := db.Select(&books, "SELECT * FROM books WHERE book_name LIKE $1", 
				"%" + bookname + "%")
	if err != nil {
		slog.Error("Problem with reading from db", "Error", err)
		return books, errors.New("Problem with reading from DB")
	}
	return books, nil
}