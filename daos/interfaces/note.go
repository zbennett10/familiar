package interfaces

import (
	"dungeon-comrade/models"

	"github.com/gocql/gocql"
)

type NoteDao interface {
	Create(user *models.CreateNoteRequest) error
	GetNotesByUser(userId gocql.UUID) ([]models.Note, error)
	// update(u *models.User) error
	// delete(id int) error
}
