package middleware

import (
	"context"
	"github.com/ganeshaditya1/GoBackend/Aggsvc/internal/authsvc"
	"github.com/ganeshaditya1/GoBackend/Aggsvc/internal/requestmodels"
	"net/http"
	"strings"
)



// returns token and true if present and well-formed ("Bearer <token>")
func BearerTokenFromRequest(r *http.Request) (string, bool) {
    auth := r.Header.Get("Authorization")
    if auth == "" {
        return "", false
    }
    // Typical form: "Bearer <token>"
    const prefix = "Bearer "
    if !strings.HasPrefix(auth, prefix) {
        return "", false
    }
    token := strings.TrimSpace(auth[len(prefix):])
    if token == "" {
        return "", false
    }
    return token, true
}

func NewBearerTokenAuth(authserv authsvc.Server) requestmodels.StrictMiddlewareFunc {
	var Authorizer = requestmodels.StrictMiddlewareFunc(func(handler requestmodels.StrictHandlerFunc, 
												endpoint string) requestmodels.StrictHandlerFunc{
		curriedHandler := func(ctx context.Context, 
						w http.ResponseWriter, 
						r *http.Request, 
						request interface{}) (interface{}, error) {
					token, ok := BearerTokenFromRequest(r)
					if !ok {
						http.Error(w, "No Authorization token found.", http.StatusForbidden)
						return nil, nil
					}

					isValid, isAdmin, err := authserv.ValidateToken(token)
					if !isValid {
						http.Error(w, "Not a valid token", http.StatusForbidden)
						return nil, nil
					}

					if err != nil {
						http.Error(w, "Unable to contact Authentication service.", http.StatusInternalServerError)
						return nil, nil
					}

					if !isAdmin && endpoint != "SearchBooks" {
						http.Error(w, "Not authorized to perform this request.", http.StatusForbidden)
					}

					return handler(ctx, w, r, request)
		}
		return requestmodels.StrictHandlerFunc(curriedHandler)
	})
	return Authorizer
}