package utils

import (
	"github.com/andrelopes-code/simple-picpay-golang/entity"
	"gorm.io/gorm"
)

// This function create an initial data just for demonstration
func SetInitialTestData(db *gorm.DB) {
	var err error
	cleanDatabase(db)

	createUser := func(name string, email string, document string, password string, balance int64, utype entity.UserType) {
		user := entity.User{
			FullName:       name,
			Document:       document,
			Email:          email,
			Password:       password,
			Type:           utype,
			BalanceInCents: balance,
		}
		err = db.Create(&user).Error
		if err != nil {
			panic(err)
		}
	}
	createUser("John", "j@a.com", "47857490273", "123456", 530, entity.CommonUser)
	createUser("Jane", "j@b.com", "00.623.904/0001-73", "123456", 5000, entity.MerchantUser)
	createUser("Mary", "m@c.com", "00.512.234/0001-23", "123456", 10000, entity.MerchantUser)
	createUser("Mark", "m@d.com", "34589673843", "123456", 2800, entity.CommonUser)
	createUser("Mike", "m@e.com", "38967580283", "123456", 10, entity.CommonUser)
}

func cleanDatabase(db *gorm.DB) {
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
}
