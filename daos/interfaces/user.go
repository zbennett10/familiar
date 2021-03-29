package interfaces

import "dungeon-comrade/models"

type UserDao interface {
	Create(user *models.CreateUserRequest) error
	GetByEmail(email string) (models.User, error)
	GetAll() ([]models.User, error)
	// update(u *models.User) error
	// delete(id int) error
}
