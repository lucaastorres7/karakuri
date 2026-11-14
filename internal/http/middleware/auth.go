package middleware

import (
	"net/http"

	"github.com/lucaastorres7/karakuri/internal/config"
	"github.com/lucaastorres7/karakuri/pkg"
)

// A middleware is any function that accepts an http.Handler and returns a new http.Handler

func RequireAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token, err := pkg.ExtractBearerToken(r)
		if err != nil {
			pkg.WriteJSON(w, http.StatusUnauthorized, map[string]string{
				"error": err.Error(),
			})
			return
		}
		
		claims, err := pkg.ValidateToken(token, config.Env.TokenSecret)
		if err != nil {
			pkg.WriteJSON(w, http.StatusUnauthorized, map[string]string{
				"error": err.Error(),
			})
			return 
		}

		ctx := pkg.SetUserClaims(r.Context(), claims)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

