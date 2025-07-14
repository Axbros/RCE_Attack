package main

import (
	"awesomeProject/handler"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/run", handler.Run)
	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
