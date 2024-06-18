package cfg

import (
	"os"
)

type AppSettings struct {
	DBName  string
	DBPath  string
	AuthURL string
}

func getSettings() *AppSettings {
	settings := &AppSettings{
		DBName:  os.Getenv("DB_NAME"),
		DBPath:  os.Getenv("DB_PATH"),
		AuthURL: os.Getenv("AUTH_URL"),
	}

	settings.validate()

	return settings
}

func (s *AppSettings) validate() {
	if s.DBPath == "" {
		panic("DB_PATH is required")
	}

	if s.AuthURL == "" {
		panic("AUTH_URL is required")
	}

	if s.DBName == "" {
		panic("DB_NAME is required")
	}
}
