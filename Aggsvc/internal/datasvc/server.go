package datasvc

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	. "github.com/ganeshaditya1/GoBackend/Aggsvc/internal/requestmodels"
	"log/slog"
	"net/http"
	"net/url"
	"time"
) 

type Server struct {
	portno int
}

func NewServer(portno int) Server {
	return Server {
		portno: portno,
	}
}

func (serv Server) CreateBook(request CreateBookRequestObject) error {
	reqBody, _ := json.Marshal(request.Body)
	req, _ := http.NewRequest("POST", 
							  fmt.Sprintf("http://localhost:%d/books", serv.portno), 
							  bytes.NewReader(reqBody))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Origin", "http://localhost")

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)

	if err != nil || resp.StatusCode != http.StatusOK {
		slog.Error("Failed to call Datasvc.", "Svc error", err)
		return err
	}
	return nil
}

func (serv Server) RemoveBook(request RemoveBookRequestObject) error {	
	req, err := http.NewRequest("DELETE", 
							  fmt.Sprintf("http://localhost:%d/books/%d", 
							  serv.portno, request.Id),
							  bytes.NewReader([]byte{}))
	req.Header.Set("Origin", "http://localhost")

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)

	if err != nil || resp.StatusCode != http.StatusOK {
		slog.Error("Failed to call Datasvc.", "Svc error", err)
		return err
	}
	return nil
}

func (serv Server) SearchBooks(request SearchBooksRequestObject) (BookCollection, error) {	
	req, err := http.NewRequest("GET", 
							  fmt.Sprintf("http://localhost:%d/books/search?%s", 
							  serv.portno, url.QueryEscape(request.Params.Name)),
							  bytes.NewReader([]byte{}))
	req.Header.Set("Origin", "http://localhost")

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)

	if err != nil || resp.StatusCode != http.StatusOK {
		slog.Error("Failed to call Datasvc.", "Svc error", err)
		return BookCollection{}, err
	}

    defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		slog.Error("Server returned non-200 response code.", 
				   "status code", resp.StatusCode)
		return BookCollection{}, errors.New("Server returned error response.")
	} else {
		var decodedResp = BookCollection{}
		err := json.NewDecoder(resp.Body).Decode(&decodedResp)
		if err != nil {
			slog.Error("Decoding Authsvc response resulted in an error.", "Authsvc error", err)
			return decodedResp, err
		}
		return decodedResp, nil
	}
	return BookCollection{}, nil
}