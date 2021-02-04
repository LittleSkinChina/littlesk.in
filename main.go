package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/LittleSkinChina/littlesk.in/handler"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "80"
	}

	http.ListenAndServe(fmt.Sprintf(":%s", port), http.HandlerFunc(handler.Handler))
}
