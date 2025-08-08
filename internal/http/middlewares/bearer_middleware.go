package middlewares

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/BrunoPolaski/go-rest-err/rest_err"
	internal_jwt "github.com/BrunoPolaski/projects-template/internal/infra/thirdparty/jwt"
	"github.com/BrunoPolaski/projects-template/internal/infra/thirdparty/logger"
	"github.com/golang-jwt/jwt/v5"
)

func BearerAuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		encoder := json.NewEncoder(w)
		header := r.Header.Get("Authorization")

		jwtAdapter := internal_jwt.NewJWTAdapter()
		token, restErr := jwtAdapter.TrimPrefix(header)
		if restErr != nil {
			w.WriteHeader(restErr.Code)
			encoder.Encode(restErr)
			return
		}

		parsedToken, restErr := jwtAdapter.ParseToken(token)
		if restErr != nil {
			w.WriteHeader(restErr.Code)
			encoder.Encode(restErr)
			return
		}

		claims, ok := parsedToken.Claims.(jwt.MapClaims)
		if !ok || !parsedToken.Valid {
			logger.Error("Invalid token claims or token is not valid", nil)

			restErr := rest_err.NewUnauthorizedError("invalid token")
			w.WriteHeader(restErr.Code)
			encoder.Encode(restErr)
			return
		}

		exp, err := claims.GetExpirationTime()
		if err != nil {
			logger.Error("Failed to get expiration time from claims: %s", err)

			restErr := rest_err.NewUnauthorizedError("invalid token")
			w.WriteHeader(restErr.Code)
			encoder.Encode(restErr)
			return
		}

		if exp.Before(time.Now()) {
			restErr := rest_err.NewUnauthorizedError("token expired")
			w.WriteHeader(restErr.Code)
			encoder.Encode(restErr)
			return
		}

		next.ServeHTTP(w, r)
	})
}
