package session

import "github.com/gorilla/securecookie"

// SessionIDGenerator generate encoded session id
var SessionIDGenerator *securecookie.SecureCookie

// Init initialize SetCookie Instance
func Init(hashKeyBase string, blockKeyBase string) {
	hashKey := []byte(hashKeyBase)
	blockKey := []byte(blockKeyBase)
	SessionIDGenerator = securecookie.New(hashKey, blockKey)
}
