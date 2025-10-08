package api

import (
	"context"
	"errors"
	"github.com/ganeshaditya1/GoBackend/Datasvc/internal/datalayer"
	"log/slog"
)

type Server struct {
}

func NewServer() Server {
	return Server{}
}

func (serv Server) CreateBook(ctx context.Context, 
						request CreateBookRequestObject) (CreateBookResponseObject, error) {
    book := datalayer.NewBook(request.Body.BookName, request.Body.BookDescription)
    err := datalayer.CreateBook(book)
	if err != nil {
		slog.Debug("Encountered an issue while inserting data into DB.", "DB error", err)
		return nil, errors.New("Internal server error.")
	}
	return CreateBook201JSONResponse{}, nil
}

func (serv Server) RemoveBook(ctx context.Context, 
	request RemoveBookRequestObject) (RemoveBookResponseObject, error) {
	err := datalayer.RemoveBook(int(request.Id))
	if err != nil {
		slog.Error("Encountered an issue while removing record from DB.", "DB error", err)
		return nil, errors.New("Internal server error.")
	}
	return RemoveBook204Response{}, nil
}

func (serv Server) SearchBooks(ctx context.Context, 
	request SearchBooksRequestObject) (SearchBooksResponseObject, error) {

	books, err := datalayer.SearchBooks(request.Params.Name)
	if err != nil {
		slog.Debug("Encountered an issue while reading from DB.", "DB error", err)
		return nil, errors.New("Internal server error.")
	}
	bookCollection := make([]Book, 0, len(books))
	for _, book := range books {
		respBook := Book{
			Bookid: int64(book.ID),
			Bookname: book.Name,
			Bookdescription: book.Description,
		}
		bookCollection = append(bookCollection, respBook)
	}	
	return SearchBooks200JSONResponse(BookCollection{
		Items: &bookCollection,
	}), nil
}