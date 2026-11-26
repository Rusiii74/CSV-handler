package main

import (
	"fmt"
	"log"
	"net/http"

	"dfs/handler"
)

func main() {
	fmt.Println("CSV process started...")

	mux := http.NewServeMux()
	mux.HandleFunc("GET /hello", handler.HelloHandler)
	mux.HandleFunc("GET /health", handler.HealthHandler)
	mux.HandleFunc("POST /upload", handler.UploadHandler)

	server := &http.Server{ // создаем сервер
		Addr:    ":8080", // указываем какой порт будет слушать сервер
		Handler: mux,     // указывает какой обраотчик будет использовать маршрутизатор
	}

	log.Fatal(server.ListenAndServe()) // запуская сервер, если порт сервера будет занят то мы закрываем программу

}
