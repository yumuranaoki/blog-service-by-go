package router

import (
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func runRouter() {
	router := mux.NewRouter()
	router.HandleFunc("/login", loginHandler)
	router.HandleFunc("/signup", signupHandler)

	server := &http.Server{
		Addr:    os.Getenv("HOST") + ":" + os.Getenv("PORT"),
		Handler: router,
	}
	server.ListenAndServe()
}

func loginHandler(w http.ResponseWriter, r *http.Request) {

}

func signupHandler(w http.ResponseWriter, r *http.Request) {

}
