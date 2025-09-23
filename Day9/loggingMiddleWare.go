// Logs every incoming request.
package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/", simpleRoute)

	loggedmux := loggingMiddleware(mux)

	fmt.Println("Server running at http://localhost:8834")
	http.ListenAndServe(":8834", loggedmux)
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		log.Printf("%s %s %s", r.Method, r.RequestURI, r.RemoteAddr)
		next.ServeHTTP(w, r)
		log.Printf("Completed in %v", time.Since(start))
	})
}

func simpleRoute(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello Handler")
}
