package cfg

import "os"

type AppSettings struct {
	DBPath  string
	AuthURL string
}

func getSettings() *AppSettings {
	settings := &AppSettings{
		DBPath:  os.Getenv("DB_PATH"),
		AuthURL: os.Getenv("AUTH_URL"),
	}
	return settings
}
