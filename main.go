package main

import (
	"fmt"

	_ "github.com/lib/pq"
	"github.com/yumuranaoki/blog-service-by-go/cacheserver"
	db "github.com/yumuranaoki/blog-service-by-go/database"
	"github.com/yumuranaoki/blog-service-by-go/router"
	"github.com/yumuranaoki/blog-service-by-go/session"
)

func main() {
	fmt.Printf("environmental variables is")
	router.RunRouter()
	db.Init()
	session.Init("hashKeyBase", "blockKeyBase")
	cacheserver.Init()
}
