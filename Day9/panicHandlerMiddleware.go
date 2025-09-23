package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/panic", panicHandler)

	recoveredMux := recoverMiddleware(mux)

	fmt.Println("Server running at http://localhost:8834...")
	http.ListenAndServe(":8834", recoveredMux)
}

func recoverMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if err := recover(); err != nil {
			log.Println("Recover from panic...", err)
			http.Error(w, "Internal server error", http.StatusInternalServerError)

			//Server didn't crash, it just logged out panic
		}
		next.ServeHTTP(w, r)
	})
}

// Sample panic message
func panicHandler(w http.ResponseWriter, r *http.Request) {
	panic("Something went wrong:(")
}
