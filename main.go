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
	isFinished := make(chan bool)
	session.Init("hashKeyBase+++++", "blockKeyBase++++")
	go db.Init(isFinished)
	go cacheserver.Init(isFinished)
	go router.RunRouter(isFinished)
	fmt.Printf("waiting for server shut down\n")
	<-isFinished
	fmt.Printf("server finished\n")
}
