package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/taaaaho/go-bot/handler"
)

func main() {
	// Create handler
	handler := handler.NewHandler()

	// Start server
	fmt.Println("START SERVER:", os.Getenv("PORT"))
	if err := http.ListenAndServe(":"+os.Getenv("PORT"), handler); err != nil {
		log.Fatal(err)
	}
}
