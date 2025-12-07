package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte("Hello World"))

		if err != nil {
			panic(err)
		}
	})

	serverAddress := fmt.Sprintf(":%s", 8080)

	server := &http.Server{
		Addr:    serverAddress,
		Handler: nil,
	}

	err := server.ListenAndServe()

	if err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
