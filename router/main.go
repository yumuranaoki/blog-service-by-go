package router

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// RunRouter start router
func RunRouter(isFinished chan bool) {
	fmt.Printf("running router\n")
	router := mux.NewRouter()
	router.HandleFunc("/login", LoginHandler).Methods("POST")
	router.HandleFunc("/signup", SignupHandler).Methods("POST")

	server := &http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: router,
	}
	server.ListenAndServe()
	fmt.Printf("server shut down\n")
	isFinished <- true
	close(isFinished)
}
