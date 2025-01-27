package main

import (
	"Communication/internal/handler"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/api/v1/communication", handler.CommunicationHandler)
	log.Println("Server started on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
