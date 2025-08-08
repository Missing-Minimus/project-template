package main

import (
	"log"
	"net/http"
	"time"

	"github.com/Missing-Minimus/projects-template/internal/http/routes"
	"github.com/Missing-Minimus/projects-template/internal/infra/thirdparty/logger"
	"github.com/joho/godotenv"
)

func main() {
	location, err := time.LoadLocation("America/Sao_Paulo")
	if err != nil {
		log.Fatalf("Error loading location: %v", err)
	}
	time.Local = location

	err = godotenv.Overload(".env")
	if err != nil {
		log.Fatalf("Error loading .env file. Error: %s", err)
	}

	logger.Init()
	logger.Info("Starting application")

	mux := http.NewServeMux()
	routes.InitRoutes(mux)
	log.Fatal(http.ListenAndServe(":8080", mux))
}
