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

func setInitialTestData(db *gorm.DB) error {
	err := db.Exec("DELETE FROM users").Error
	if err != nil {
		panic("Error deleting data from users table: " + err.Error())
	}

	err = db.Exec("DELETE FROM sqlite_sequence WHERE name='users'").Error
	if err != nil {
		panic("Error resetting sequence for users table: " + err.Error())
	}

	err = db.Exec("DELETE FROM transactions").Error
	if err != nil {
		panic("Error deleting data from transactions table: " + err.Error())
	}

	err = db.Exec("DELETE FROM sqlite_sequence WHERE name='transactions'").Error
	if err != nil {
		panic("Error resetting sequence for transactions table: " + err.Error())
	}

	merchantUser := entity.User{
		FullName:       "Alvaro Morata",
		DocumentID:     "00.623.904/0001-73",
		Email:          "pR3z7asd@example.com",
		Password:       "123456",
		Type:           entity.CommonUser,
		BalanceInCents: 2276,
	}

	err = db.Create(&merchantUser).Error
	if err != nil {
		return err
	}

	individualUser := entity.User{
		FullName:       "Andres Felipe",
		DocumentID:     "12345678912",
		Email:          "sfgjcv43f@example.com",
		Password:       "123456",
		Type:           entity.CommonUser,
		BalanceInCents: 3555,
	}

	err = db.Create(&individualUser).Error
	if err != nil {
		return err
	}

	return nil
}
