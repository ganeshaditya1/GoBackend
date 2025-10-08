package middleware

import (
	"context"
	"github.com/ganeshaditya1/GoBackend/Datasvc/internal/api"
	"github.com/go-playground/validator/v10"
	"net/http"	
)

var Validator = api.StrictMiddlewareFunc(func(handler api.StrictHandlerFunc, 
											  endpoint string) api.StrictHandlerFunc{
	curriedHandler := func(ctx context.Context, 
					w http.ResponseWriter, 
					r *http.Request, 
					request interface{}) (interface{}, error) {
				validate := validator.New(validator.WithRequiredStructEnabled())
				errs := validate.Struct(request)
				if errs != nil {
					http.Error(w, errs.Error(), http.StatusBadRequest)
					return nil, nil
				}
				return handler(ctx, w, r, request)
	}
	return api.StrictHandlerFunc(curriedHandler)
})