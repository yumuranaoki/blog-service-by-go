package session

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/securecookie"
)

// SessionIDGenerator generate encoded session id
var SessionIDGenerator *securecookie.SecureCookie

// Init initialize SetCookie Instance
func Init(hashKeyBase string, blockKeyBase string) {
	hashKey := []byte(hashKeyBase)
	blockKey := []byte(blockKeyBase)
	SessionIDGenerator = securecookie.New(hashKey, blockKey)
}

// ReadSessionIDFromCookie reads session id from cookie and return it
func ReadSessionIDFromCookie(w http.ResponseWriter, r *http.Request) string {
	cookie, err := r.Cookie("session_id")
	if err != nil {
		log.Fatal(err)
	}
	sessionData := make(map[string]int)
	if err := SessionIDGenerator.Decode("sessionID", cookie.Value, &sessionData); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("session id: %s\n", string(sessionData["sessionID"]))

	return string(sessionData["sessionID"])
}
