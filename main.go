package main

import (
	"fmt"
	"os"

	"github.com/yumuranaoki/blog-serverice-by-go/router"
)

func main() {
	fmt.Printf("environmental variables is %s\n", os.Getenv("HOST"))
	router.RunRouter()
}
