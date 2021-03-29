package cassandra

import (
	"dungeon-comrade/models"
	"fmt"
)

type UserImpl struct {
}

func (dao UserImpl) Create(user *models.CreateUserRequest) error {
	session := get()

	insertionError := session.Query(
		"INSERT INTO user (id, join_date, firstname, lastname, email) VALUES (uuid(), now(), ?, ?, ?) IF NOT EXISTS",
		user.FirstName, user.LastName, user.Email,
	).Exec()

	if insertionError != nil {
		fmt.Printf("Create User Error: %s\n", insertionError)
		return insertionError
	}

	defer session.Close()

	return nil
}

func (dao UserImpl) GetByEmail(email string) (models.User, error) {
	session := get()

	var user models.User

	err := session.Query(
		"SELECT id, email, join_date, firstname, lastname FROM user WHERE email = ? LIMIT 1 ALLOW FILTERING;", email).Scan(
		&user.ID, &user.Email, &user.JoinDate, &user.FirstName, &user.LastName)

	if err != nil {
		fmt.Printf("Fetch User Error: %s\n", err)
	}

	defer session.Close()

	return user, err
}

func (dao UserImpl) GetAll() ([]models.User, error) {
	session := get()

	var users []models.User

	err := session.Query("SELECT * FROM user;").Scan(&users)

	if err != nil {
		fmt.Printf("Fetch All Users Error: %s\n", err)
	}

	defer session.Close()

	return users, err
}
