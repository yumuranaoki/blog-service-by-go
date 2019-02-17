package router

import (
	"fmt"
	"net/http"

	"github.com/yumuranaoki/blog-service-by-go/session"
)

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("root handler")
	sessionID := session.ReadSessionIDFromCookie(w, r)
	w.Write([]byte(sessionID))
}
