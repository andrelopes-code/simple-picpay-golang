package repository

import (
	"github.com/andrelopes-code/simple-picpay-golang/cfg"
	"gorm.io/gorm"
)

var (
	log = cfg.GetLogger("repository")
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db}
}

func (r *Repository) GetDB() *gorm.DB {
	return r.db
}
