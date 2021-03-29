package cassandra

import (
	"dungeon-comrade/models"

	"github.com/gocql/gocql"
)

type NoteImpl struct {
}

func (dao NoteImpl) Create(user *models.CreateNoteRequest) error {
	session := get()

	// insertionError := session.Query(
	// 	"INSERT INTO user (id, join_date, firstname, lastname, email) VALUES (uuid(), now(), ?, ?, ?) IF NOT EXISTS",
	// 	user.FirstName, user.LastName, user.Email,
	// ).Exec()

	// if insertionError != nil {
	// 	fmt.Printf("Create User Error: %s\n", insertionError)
	// 	return insertionError
	// }

	defer session.Close()

	return nil
}

func (dao NoteImpl) GetNotesByUser(userId gocql.UUID) ([]models.Note, error) {
	session := get()

	defer session.Close()

	var placeholder []models.Note

	return placeholder, nil
}
