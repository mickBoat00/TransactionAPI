// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0

package database

import (
	"time"

	"github.com/google/uuid"
)

type Currency struct {
	ID        uuid.UUID
	Name      string
	Code      string
	Createdat time.Time
	Updatedat time.Time
}

type User struct {
	ID        uuid.UUID
	Email     string
	Password  string
	Createdat time.Time
	Updatedat time.Time
}
