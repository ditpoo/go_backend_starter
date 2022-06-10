package main

import (
	// "os"
	"net/http"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	// "starter/common"
	"starter/handlers"
)

func main() {
	log.SetFormatter(&log.TextFormatter{
		DisableColors: true,
		FullTimestamp: true,
	})
	router := mux.NewRouter()

	// settings := common.GetSettings(os.Getenv("DHANIQ_ENV"))

	router.HandleFunc("/health", handlers.HealthHandler).Methods("GET", "OPTIONS")

	http.Handle("/", router)

	http.ListenAndServe(":5000", router)
}