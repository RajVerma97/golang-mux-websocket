package main

import (
	"log"
	"net/http"
)

func main() {

	setupAPI()

	log.Println("Listening on Port 8080")
	http.ListenAndServe(":8080", nil)

}

func setupAPI() {
	manager := NewManager()
	http.Handle("/", http.FileServer(http.Dir("./frontend")))
	http.HandleFunc("/ws", manager.serveWS)

}
