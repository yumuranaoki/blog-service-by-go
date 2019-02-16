package router

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/yumuranaoki/blog-service-by-go/cacheserver"
	db "github.com/yumuranaoki/blog-service-by-go/database"
	"github.com/yumuranaoki/blog-service-by-go/model"
	"github.com/yumuranaoki/blog-service-by-go/session"
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

	// check if data is valid
	if true := db.DB.NewRecord(user); true {
		db.DB.Create(&user)
	}

	id := user.ID
	sessionData := map[string]uint{
		"sessionID": id,
	}
	sessionToken, err := session.SessionIDGenerator.Encode("sessionID", sessionData)
	if err != nil {
		log.Fatal(err)
	}

	// 0 expiration time means the key has no expiration time
	err = cacheserver.RedisClient.Set("session_id", id, 0).Err()
	if err != nil {
		log.Fatal(err)
	}

	cookie := &http.Cookie{
		Name:     "session_id",
		Value:    sessionToken,
		HttpOnly: true,
	}

	http.SetCookie(w, cookie)
}
