package router

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// RunRouter start router
func RunRouter() {
	fmt.Printf("running router")
	router := mux.NewRouter()
	router.HandleFunc("/login", LoginHandler).Methods("POST")
	router.HandleFunc("/signup", SignupHandler).Methods("POST")

	server := &http.Server{
		Addr:    "localhost" + ":" + "3000",
		Handler: router,
	}
	server.ListenAndServe()
}
