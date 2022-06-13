package main

import (
	"os"
	"net/http"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"

	// "starter/common"
	"starter/common"
	"starter/handlers"
)

func main() {
	log.SetFormatter(&log.TextFormatter{
		DisableColors: true,
		FullTimestamp: true,
	})
	router := mux.NewRouter()

	GetSettings := common.GetSettings(os.Getenv("DHANIQ_ENV"))

	auth0Handler := handlers.Auth0{
		Settings: GetSettings,
	}

	router.HandleFunc("/health", handlers.HealthHandler).Methods("GET", "OPTIONS")
	router.HandleFunc("/login", auth0Handler.Login).Methods("POST", "OPTIONS")

	http.Handle("/", router)

	http.ListenAndServe(":5000", router)
}