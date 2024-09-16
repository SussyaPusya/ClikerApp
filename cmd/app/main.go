package main

import (
	"clickercopro/app/handlers"
	"fmt"
	"net/http"
	"os"
	// пакеты
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	mux := http.NewServeMux()

	fsApp := http.FileServer(http.Dir("./cmd/app"))
	mux.Handle("/", fsApp)

	fsWeb := http.FileServer(http.Dir("./cmd/web"))
	mux.Handle("/web/", http.StripPrefix("/web/", fsWeb))

	mux.HandleFunc("/endpoint", handlers.HandlerClickevent) // Используем mux.HandleFunc

	fmt.Println("Server has Started")
	http.ListenAndServe(":"+port, mux)
}

//
