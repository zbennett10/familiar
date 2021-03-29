package models

import "github.com/gocql/gocql"

type CreateUserRequest struct {
	FirstName string
	LastName  string
	Email     string
}

type User struct {
	ID        gocql.UUID
	JoinDate  string
	FirstName string
	LastName  string
	Email     string
}
