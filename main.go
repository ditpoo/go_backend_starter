package main

import (
	// "os"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	// "starter/common"
	"starter/handlers"
)

func removeTrailingSlash(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		req.URL.Path = strings.TrimSuffix(req.URL.Path, "/")
		next.ServeHTTP(w, req)
	})
}

func main() {
	log.SetFormatter(&log.TextFormatter{
		DisableColors: true,
		FullTimestamp: true,
	})
	router := mux.NewRouter()

	// settings := common.GetSettings(os.Getenv("DHANIQ_ENV"))

	router.HandleFunc("/health", handlers.HealthHandler).Methods("GET", "OPTIONS")


	http.Handle("/", router)

	http.ListenAndServe(":5000", removeTrailingSlash(router))
}