package middlewares

import (
	"context"
	"net/http"

	"danieljonguitud.com/aws-events-go/auth"
)

func AuthHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")

		if token == "" {
			http.Error(w, "Provide a token", http.StatusBadRequest)
			return
		}

		userId, err := auth.VerifyToken(token)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		ctx := context.WithValue(r.Context(), "userId", userId)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
