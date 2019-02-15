package router

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/yumuranaoki/blog-serverice-by-go/model"
)

// SignupHandler handle request to signup for signup
func SignupHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("post is done,\n")
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Fatal(err)
	}
	var user model.User
	if err := json.Unmarshal(body, &user); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Fatal(err)
	}
	fmt.Printf("name: %s\nemail: %s\npassword: %s\n", user.Name, user.Email, user.Password)
}
