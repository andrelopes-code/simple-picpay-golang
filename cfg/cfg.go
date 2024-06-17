package cfg

import (
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

var (
	Settings *AppSettings
	logger   *Logger
	db       *gorm.DB
)

func Init() {
	// Initialize Logger
	logger = GetLogger("cfg")
	logger.Info("Initializing config...")

	// Load .env file
	err := godotenv.Load(".env")
	if err != nil {
		logger.Fatal("Error loading .env file: " + err.Error())
	}

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
