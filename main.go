package main

import (
	"awesomeProject/handler"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/run", handler.Run)
	log.Println("Server started on :613")
	log.Fatal(http.ListenAndServe(":613", nil))
}
