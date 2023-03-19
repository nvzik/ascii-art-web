package main

import (
	"ascii-art-web/handlers"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handlers.HomeHandler)
	mux.HandleFunc("/ascii-art", handlers.PostHandler)
	mux.Handle("/favicon/", http.StripPrefix("/favicon/", http.FileServer(http.Dir("favicon"))))
	log.Println("Запуск веб-сервера на http://localhost:8000/ ")
	err := http.ListenAndServe(":8000", mux)
	if err != nil {
		log.Fatal(err)
	}
}
