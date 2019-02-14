package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Printf("environmental variables is %s\n", os.Getenv("HOST"))
}
