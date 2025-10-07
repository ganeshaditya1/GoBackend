package api

import (
	"context"
	"errors"
	"github.com/ganeshaditya1/GoBackend/Authsvc/internal/datalayer"
	"github.com/ganeshaditya1/GoBackend/Authsvc/internal/util"
	"log/slog"
	"strings"
	"strconv"

	openapi_types "github.com/oapi-codegen/runtime/types"
)

type Server struct {
	tokens map[string]string
	jwtHandler util.JWTHandler
}

func NewServer(jwtHandler util.JWTHandler) Server {
	serverObj := Server{jwtHandler: jwtHandler}
	serverObj.tokens = make(map[string]string)
	return serverObj
}



func (serv Server) LoginUser(ctx context.Context, 
						request LoginUserRequestObject) (LoginUserResponseObject, error) {
    user, err := datalayer.LoginUser(request.Body.Username, request.Body.Password)
	if err != nil {
		if strings.Contains(err.Error(), "Incorrect password") {
			msg := "Incorrect credentials"
			resp := UnauthorizedJSONResponse{
				Message: &msg,
			}
			return LoginUser401JSONResponse{resp}, nil
		} else {
			slog.Debug("Encountered an issue while inserting data into DB.", "DB error", err)
			return nil, errors.New("Internal server error.")
		}
	}
	jwtUserClaims := util.NewJWTUserClaims(
		user.Userid,
		user.Username,
		user.EmailAddress,
		user.IsAdmin,
	)
	Token := serv.jwtHandler.GenerateToken(jwtUserClaims)
	Message := "Logged In" 
	resp := LoginSuccessfulJSONResponse{
		Message: &Message,
		Token: &Token,
	}

	serv.tokens[request.Body.Username] = Token
	return LoginUser200JSONResponse{resp}, nil
}

func (Server) CreateUser(ctx context.Context, 
						 request CreateUserRequestObject) (CreateUserResponseObject, error){
    user := datalayer.NewUser(request.Body.Username, 
							  string(request.Body.Email),
							  request.Body.Password,
							  request.Body.Age)

	err := datalayer.CreateUser(user)
	if err != nil {
		if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
			msg := "User account by that details already exists"
			resp := ConflictJSONResponse {
				Message: &msg,
			}
			return CreateUser409JSONResponse{resp}, nil
		} else {
			slog.Debug("Encountered an issue while inserting data into DB.", "DB error", err)
			return nil, errors.New("Internal server error.")
		}
		
	} else {
		msg := "User account created."
		resp := UserCreatedJSONResponse {
			Message: &msg,
		}
		return CreateUser201JSONResponse{resp}, nil
	}
}



func (serv Server) DecodeUserToken(ctx context.Context, request DecodeUserTokenRequestObject) (DecodeUserTokenResponseObject, error) {
	token := request.Body.Token

	userClaims, errs := serv.jwtHandler.DecodeToken(token)
	if errs != nil {
		slog.Error("Token failed to parse", errs)
		return InvalidTokenResponse{}, nil		
	}

	cachedToken, ok := serv.tokens[userClaims.Username]

	if (!ok || cachedToken != token) {
		slog.Error("Either user never logged in or supplied token is invalid.")
		return InvalidTokenResponse{}, nil
	}

	userid, _ := strconv.Atoi(userClaims.Subject)
	userEmail := openapi_types.Email(userClaims.Email)
	decodedToken := ValidTokenJSONResponse {
		Emailid: &userEmail,
		Isadmin: &userClaims.IsAdmin,
		Userid: &userid,
		Username: &userClaims.Username,
	}
	return DecodeUserToken200JSONResponse{decodedToken}, nil	
}