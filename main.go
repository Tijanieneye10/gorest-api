package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/tijanieneye10/restapi/internal/config"
	"github.com/tijanieneye10/restapi/internal/handlers"
	"github.com/tijanieneye10/restapi/internal/routes"
	"github.com/tijanieneye10/restapi/internal/store"
)

func main() {

	configValues, err := config.LoadConfig()

	//connect db
	db := config.ConnectDB(configValues.DatabaseURL)

	defer db.Close()

	if err != nil {
		log.Fatal(err)
	}
	mux := http.NewServeMux()

	queries := store.New(db)

	handler := handlers.NewHandler(db, queries)

	routes.SetupRoutes(mux, handler)

	serverAddress := fmt.Sprintf(":%s", configValues.ServerPort)

	server := &http.Server{
		Addr:    serverAddress,
		Handler: mux,
	}

	err = server.ListenAndServe()

	if err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
