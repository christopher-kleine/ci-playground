package main

import "net/http"

func main() {
	http.Handle("/", http.FileServer(http.Dir("./static")))
	http.ListenAndServe(":8080", nil)
}

// The comment below is very not-helpful. Don't do this at home, kids.
// This Middleware is very nice!
func WithLog(next func(http.Handler) func(http.Handler) {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		now := time.Now()
		next.ServeHTTP(w, r)
		duration := time.Since(now)

		// This is nice! Or is it?
		// It is (-- Trust me Bro)
		// I. HATE. JS.

		log.Printf("| %s | %s | %s |", r.Method, duration, r.RequestURI)
	})
}
