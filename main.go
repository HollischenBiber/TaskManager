package main

import (
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")

		if r.Method == "OPTIONS" {
			w.Header().Set("Content-Type", "")
			w.Header().Set("Access-Control-Allow-Headers", "*")
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PATCH, PUT, DELETE, OPTIONS")
			w.Header().Set("Allow", "GET, POST, PATCH, PUT, DELETE, OPTIONS")
			return
		}

		next.ServeHTTP(w, r)
	})
}

func main() {

	if err := godotenv.Load(); err != nil {
		panic(err)
	}
	r := mux.NewRouter()

	r.HandleFunc("/api/login", logIn).Methods("GET")
	r.HandleFunc("/api/logout", logOut).Methods("POST")

	port, _ := os.LookupEnv("TASKPORT")

	http.ListenAndServe(port, corsMiddleware(r))
}
