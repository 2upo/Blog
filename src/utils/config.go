package utils

import (
	"os"
	"sync"
)

var (
	configuration Configuration
	onceConfig    sync.Once
)

type Configuration struct {
	Dsn       string
	DBName    string
	ServerDsn string
}

func Config() *Configuration {
	onceConfig.Do(func() {
		configuration = Configuration{
			Dsn:       os.Getenv("MONGO_DSN"),
			DBName:    os.Getenv("MONGO_DB"),
			ServerDsn: os.Getenv("Server_DSN"),
		}
	})
	return &configuration
}
