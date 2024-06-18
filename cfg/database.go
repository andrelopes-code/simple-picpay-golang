package cfg

import (
	"os"

	"github.com/andrelopes-code/simple-picpay-golang/entity"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func initializeDatabase() (*gorm.DB, error) {
	log := GetLogger("db")
	log.Info("Initializing database...")

	dbPath := Settings.DBPath
	dbFile := dbPath + "/" + Settings.DBName

	_, err := os.Stat(dbFile)
	if os.IsNotExist(err) {
		log.Debug("Database not found, creating...")

		err := os.MkdirAll(dbPath, 0755)
		if err != nil {
			return nil, err
		}

		file, err := os.Create(dbFile)
		if err != nil {
			return nil, err
		}
		file.Close()
	}

	// Try connect to database
	db, err := gorm.Open(sqlite.Open(dbFile), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	models := []any{
		&entity.Transaction{},
		&entity.User{},
	}

	// Auto migrate database with gorm
	err = db.AutoMigrate(models...)
	if err != nil {
		return nil, err
	}

	return db, nil
}
