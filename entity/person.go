package entity

import "github.com/google/uuid"

type Person struct {
	ID       uuid.UUID `db:"id"`
	Name     string    `db:"name"`
	Email    string    `db:"email"`
	Password string    `db:"password"`
	Age      int       `db:"age"`
}
