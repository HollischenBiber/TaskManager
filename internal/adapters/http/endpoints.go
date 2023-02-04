package http

import (
	"net/http"

	"github.com/HollischenBiber/TaskManager.git/internal/domain/auth"
	"github.com/gorilla/mux"
)

func InitRouter() http.Handler {

	r := mux.NewRouter()

	r.HandleFunc("/api/login", logIn).Methods("GET")
	r.HandleFunc("/api/logout", logOut).Methods("POST")

	return r
}

func logIn(w http.ResponseWriter, r *http.Request) {

	login, password, ok := r.BasicAuth()

	if !ok {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	_, err := auth.Service.SignIn(r.Context(), login, password)

	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		return
	}

}
