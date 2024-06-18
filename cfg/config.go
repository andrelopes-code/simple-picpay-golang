package cfg

import (
	"gorm.io/gorm"
)

var (
	Settings *AppSettings
	logger   *Logger
	db       *gorm.DB
)

func Init() {
	var err error

	// Initialize Logger
	logger = GetLogger("cfg")
	logger.Info("Initializing config...")

	// Initialize Settings
	Settings = getSettings()

	// Initialize Database configuration
	db, err = initializeDatabase()
	if err != nil {
		logger.Fatalf("Error initializing database: %s", err)
	}

	// !
	setInitialTestData(db)
}

// GetLogger returns a new logger with the given prefix
func GetLogger(prefix string) *Logger {
	return newLogger(prefix)
}

// GetDB returns the database connection
func GetDB() *gorm.DB {
	return db
}

func SetDB(ndb *gorm.DB) {
	db = ndb
}
