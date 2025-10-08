package api

import (
	"context"
	"errors"
	"github.com/ganeshaditya1/GoBackend/Aggsvc/internal/datasvc"
	. "github.com/ganeshaditya1/GoBackend/Aggsvc/internal/requestmodels"
)

type Server struct {
	datashard1, datashard2 datasvc.Server
}

func NewServer(datashard1, datashard2 datasvc.Server) Server {
	return Server{datashard1: datashard1, datashard2: datashard2}
}

func FirstInAtoM(s string) bool {
    if s == "" {
        return false
    }
    c := s[0]
    if 'A' <= c && c <= 'M' {
        return true
    }
	return false
}

func (serv Server) CreateBook(ctx context.Context, 
						request CreateBookRequestObject) (CreateBookResponseObject, error) {

	var err error
	if FirstInAtoM(request.Body.BookName) {
		err = serv.datashard1.CreateBook(request)
	} else {
		err = serv.datashard2.CreateBook(request)		
	}

	if err != nil {
		return nil, errors.New("Internal server error.")
	}
								
	return CreateBook201JSONResponse{}, nil
}

func (serv Server) RemoveBook(ctx context.Context, 
	request RemoveBookRequestObject) (RemoveBookResponseObject, error) {
	serv.datashard1.RemoveBook(request)
	serv.datashard2.RemoveBook(request)
	return RemoveBook204Response{}, nil
}

func (serv Server) SearchBooks(ctx context.Context, 
	request SearchBooksRequestObject) (SearchBooksResponseObject, error) {
	bookCollection1, _ := serv.datashard1.SearchBooks(request)
	bookCollection2, _ := serv.datashard2.SearchBooks(request)
	allBooks := append(*bookCollection1.Items, *bookCollection2.Items...)
	bookCollection1.Items = &allBooks
	return SearchBooks200JSONResponse(bookCollection1), nil
}
