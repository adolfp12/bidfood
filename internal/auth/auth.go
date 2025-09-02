package auth

import (
	"crypto/subtle"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/julienschmidt/httprouter"
)

func init() {
	godotenv.Load()
}

func APIKeyAuthMiddleware(next func(w http.ResponseWriter, r *http.Request, param httprouter.Params)) func(w http.ResponseWriter, r *http.Request, param httprouter.Params) {
	return func(w http.ResponseWriter, r *http.Request, param httprouter.Params) {
		expectedAPIKey := os.Getenv("API_KEY") // Store your key in env vars
		apiKey := r.Header.Get("X-API-Key")

		// log.Printf(">>> expectedAPIKey %s =", expectedAPIKey)
		// log.Println(">>> Check API Key", subtle.ConstantTimeCompare([]byte(apiKey), []byte(expectedAPIKey)) != 1)
		if subtle.ConstantTimeCompare([]byte(apiKey), []byte(expectedAPIKey)) != 1 {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		next(w, r, param)
	}
}
