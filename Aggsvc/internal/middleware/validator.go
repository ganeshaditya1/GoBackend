package middleware

import (
	"context"
	"github.com/ganeshaditya1/GoBackend/Aggsvc/internal/requestmodels"
	"github.com/go-playground/validator/v10"
	"net/http"	
)

var Validator = requestmodels.StrictMiddlewareFunc(func(handler requestmodels.StrictHandlerFunc, 
											  endpoint string) requestmodels.StrictHandlerFunc{
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
	return requestmodels.StrictHandlerFunc(curriedHandler)
})