package endpoints

import "net/http"

func AllowCors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Set Access-Control-Allow-Origin header to allow requests from any origin
		w.Header().Set("Access-Control-Allow-Origin", "*")

		// Call the next handler in the chain
		next.ServeHTTP(w, r)
	})
}

func AllowCorsHandler(w http.ResponseWriter, request *http.Request) {
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

	// Respond with a 200 OK status to preflight requests
	w.WriteHeader(http.StatusOK)
	w.WriteHeader(http.StatusOK)
}
