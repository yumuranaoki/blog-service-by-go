package router

import (
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

// RunRouter start router
func RunRouter() {
	router := mux.NewRouter()
	router.HandleFunc("/login", LoginHandler).Methods("POST")
	router.HandleFunc("/signup", SignupHandler).Methods("POST")

	server := &http.Server{
		Addr:    os.Getenv("HOST") + ":" + os.Getenv("PORT"),
		Handler: router,
	}
	server.ListenAndServe()
}
