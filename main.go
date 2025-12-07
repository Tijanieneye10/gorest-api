package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/tijanieneye10/restapi/internal/config"
)

func main() {

	configValues, err := config.LoadConfig()

	if err != nil {
		return
	}
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte("Hello World"))

		if err != nil {

			panic(err)
		}
	})

	serverAddress := fmt.Sprintf(":%s", configValues.ServerPort)

	server := &http.Server{
		Addr:    serverAddress,
		Handler: nil,
	}

	err := server.ListenAndServe()

	if err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
