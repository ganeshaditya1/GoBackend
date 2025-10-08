package datalayer

type Book struct {
    ID int `db:"book_id"`
    Name  string `db:"book_name"`
    Description  string `db:"book_description"`
}

func NewBook(bookname string, description string) Book {
	return Book{
		Name: bookname,
		Description: description,
	}
}