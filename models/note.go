package models

import "github.com/gocql/gocql"

type CreateNoteRequest struct {
	Content string
	Tags    []string
	UserId  gocql.UUID
}

type NoteAttachment struct {
	UploadDate string
	Data       []byte
}

type Note struct {
	ID           gocql.UUID
	UserId       gocql.UUID
	Content      string
	Tags         []string
	Attachments  []NoteAttachment
	CreateDate   string
	ModifiedDate string
}
