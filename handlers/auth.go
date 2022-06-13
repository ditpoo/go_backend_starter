package handlers

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"starter/common"

	log "github.com/sirupsen/logrus"
)

type Auth0 struct {
	Settings *common.Settings
}

type credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type loginResponse struct {
	Access_token string `json:"access_token"`
}

func (s *Auth0) Login(resp http.ResponseWriter, req *http.Request) {
	// make a post request to auth0 and get the token 

	reqBody, err := ioutil.ReadAll(req.Body)

	if err != nil {
		log.Fatalf("An Error Occured %v", err)
		// handle err
	}

	var body credentials

	err = json.Unmarshal(reqBody, &body)

	if err != nil {
		log.Fatalf("An Error Occured %v", err)
		// handel err
	}

	// encode data
	postBody, _ := json.Marshal(map[string]string{
		"grant_type": "password",
		"username": body.Username,
		"password": body.Password,
		"audience": s.Settings.Auth0EndPoint + "/api/v2/",
		"scope": "profile",
		"client_id": s.Settings.Auth0ClientId,
		"client_secret": s.Settings.Auth0ClientSecret,
		"connection": "Username-Password-Authentication",
	})

	responseBody := bytes.NewBuffer(postBody)

	url := s.Settings.Auth0EndPoint + "/oauth/token"

	response, err := http.Post(url, "application/json", responseBody)

	if err != nil {
		log.Fatalf("An Error Occured %v", err)
	}

	defer response.Body.Close()

	responseData, err := ioutil.ReadAll(response.Body)

	if err != nil {
		log.Fatalf("An Error Occured %v", err)
	}

	var auth0response loginResponse

	json.Unmarshal(responseData, &auth0response)

	resp.Header().Set("Content-Type", "application/json")
	resp.WriteHeader(http.StatusCreated)
	json.NewEncoder(resp).Encode(auth0response)
}