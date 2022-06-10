package handlers

import (
	"encoding/json"
	"net/http"
)

func HealthHandler(resp http.ResponseWriter, req *http.Request) {	
	resp.WriteHeader(http.StatusOK)
}