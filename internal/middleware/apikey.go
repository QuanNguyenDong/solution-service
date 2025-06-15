package middleware

import (
	"fmt"
	"net/http"
	"os"
)

func APIKeyMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		apiKey := r.Header.Get("X-API-Key")
		validKey := os.Getenv("API_KEY")
		fmt.Println(validKey)

		if apiKey == "" || apiKey != validKey {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		
		next.ServeHTTP(w, r)
	})
}
