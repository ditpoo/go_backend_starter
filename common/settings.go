package common

import (
	"fmt"
	"os"
)

type Settings struct {
	Env                  string
	DatabaseHost         string
	DatabaseUser         string
	DatabasePassword     string
	Auth0ClientId        string
	Auth0ClientSecret    string
}

func (s *Settings) GetDBString() string {
	return fmt.Sprintf("host=%s user=%s password=%s sslmode=disable", s.DatabaseHost, s.DatabaseUser, s.DatabasePassword)
}

func (s *Settings) IsProduction() bool {
	return s.Env == "prod"
}

func GetSettings(env string) *Settings {
	switch env {
	default:
		return &Settings{
			Env: "dev",
			DatabaseHost: os.Getenv("DATABASE_HOST"),
			DatabaseUser: os.Getenv("DATABASE_USER"),
			DatabasePassword: os.Getenv("DATABASE_PASSWORD"),
			Auth0ClientId: os.Getenv("AUTH0_CLIENT_ID"),
			Auth0ClientSecret: os.Getenv("AUTH0_CLIENT_SECRET"),
		}
	case "prod":
		return &Settings{
			Env: "prod",
			DatabaseHost: os.Getenv("DATABASE_HOST"),
			DatabaseUser: os.Getenv("DATABASE_USER"),
			DatabasePassword: os.Getenv("DATABASE_PASSWORD"),
			Auth0ClientId: os.Getenv("AUTH0_CLIENT_ID"),
			Auth0ClientSecret: os.Getenv("AUTH0_CLIENT_SECRET"),
		}
	}
}