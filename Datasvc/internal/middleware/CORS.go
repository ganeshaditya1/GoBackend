package middleware

import (
	"context"
	"github.com/ganeshaditya1/GoBackend/Datasvc/internal/api"
	"net/http"
	"net/url"
	"strings"
)

func originAllowed(origin string, allowed []string) bool {
	if origin == "" {
		return false
	}
	u, err := url.Parse(origin)
	if err != nil {
		return false
	}
	origFull := u.Scheme + "://" + u.Host // includes port if present
	for _, a := range allowed {
		if a == "*" {
			return true
		}
		// exact full-origin match (scheme must match)
		if a == origFull {
			return true
		}
		// wildcard with scheme, e.g. "https://*.example.com"
		if strings.HasPrefix(a, "http://") || strings.HasPrefix(a, "https://") {
			// replace scheme://*.example.com -> scheme:// and check suffix on origFull
			if strings.HasPrefix(a, u.Scheme+"://") && strings.HasPrefix(a, u.Scheme+"://*.") {
				if strings.HasSuffix(origFull, a[len(u.Scheme+"://*"):]) {
					return true
				}
			}
		}
		// host-only wildcard e.g. "*.example.com" or exact host "api.example.com"
		if strings.HasPrefix(a, "*.") {
			if strings.HasSuffix(u.Host, a[1:]) {
				return true
			}
		}
		if a == u.Host {
			return true
		}
	}
	return false
}

func NewCORSFilter(allowed_origins []string) api.StrictMiddlewareFunc {

	var middleware = api.StrictMiddlewareFunc(func(handler api.StrictHandlerFunc, 
											endpoint string) api.StrictHandlerFunc{
		curriedHandler := func(ctx context.Context, 
						w http.ResponseWriter, 
						r *http.Request, 
						request interface{}) (interface{}, error) {							
					
					origin := r.Header.Get("Origin")

					// Decide Access-Control-Allow-Origin
					if originAllowed(origin, allowed_origins) {
						// if allowed origins is "*" and credentials not allowed, can set "*" else echo origin
						if len(allowed_origins) == 1 && allowed_origins[0] == "*" {
							w.Header().Set("Access-Control-Allow-Origin", "*")
						} else {
							w.Header().Set("Access-Control-Allow-Origin", origin)
						}
					} else {
						// origin not allowed: respond with 403 for preflight or just continue without CORS headers
						http.Error(w, "Origin not allowed", http.StatusForbidden)
						return nil, nil
					}
					return handler(ctx, w, r, request)
		}
		return api.StrictHandlerFunc(curriedHandler)
	})
	return middleware
}