package authsvc

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
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

type DecodeTokenRequest struct {
	Token string `json:"token"`
}

type DecodeTokenResponse struct {
	IsAdmin bool `json:"isadmin"`
}

func (serv Server) ValidateToken(token string) (isValidToken bool, isAdmin bool, problem error) {
	reqBody, _ := json.Marshal(DecodeTokenRequest{Token: token})
	req, _ := http.NewRequest("POST", 
							  fmt.Sprintf("http://localhost:%d/decode-usertoken", serv.portno), 
							  bytes.NewReader(reqBody))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Origin", "http://localhost")


	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)

	if err != nil {
		problem = err
		slog.Error("Failed to call Authsvc.", "Svc error", err)
		return
	}

    defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return
	} else {
		var decodedResp = DecodeTokenResponse{}
		err := json.NewDecoder(resp.Body).Decode(&decodedResp)
		if err != nil {
			slog.Error("Decoding Authsvc response resulted in an error.", "Authsvc error", err)
			problem = err
			return
		}
			isValidToken = true
			isAdmin = decodedResp.IsAdmin
			return
	}
}