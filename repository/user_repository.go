package repository

import (
	"github.com/andrelopes-code/simple-picpay-golang/entity"
)

// FindUserByID returns a user by its ID
func (r *Repository) FindUserByID(id uint64) (*entity.User, error) {
	user := &entity.User{}
	err := r.db.First(user, id).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

// CreateUser creates a new user in the database
func (r *Repository) SaveUser(user *entity.User) error {
	return r.db.Create(user).Error
}

// UpdateUser updates a user in the database
func (r *Repository) UpdateUser(user *entity.User) error {
	return r.db.Save(user).Error
}

// DeleteUser deletes a user from the database
func (r *Repository) DeleteUser(user *entity.User) error {
	return r.db.Delete(user).Error
}
