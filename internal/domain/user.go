package domain

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type (
	User struct {
		id          uuid.UUID
		email       string
		password    string
		firstname   string
		lastname    string
		dateOfBirth time.Time
		createdAt   time.Time
		updatedAt   time.Time
	}

	UserRepository interface {
		Create(context.Context, User) (User, error)
		// FindAll(context.Context) ([]User, error)
		// Update(context.Context, User) (User, error)
		// FindById(context.Context, uuid.UUID) (User, error)
		// DeleteById(context.Context, uuid.UUID) error
	}
)

func NewUser(id uuid.UUID, email, password, firstname, lastname string, dateOfBirth, createdAt, updatedAt time.Time) User {
	return User{
		id:          id,
		email:       email,
		password:    password,
		firstname:   firstname,
		lastname:    lastname,
		dateOfBirth: dateOfBirth,
		createdAt:   createdAt,
		updatedAt:   updatedAt,
	}
}

func (u User) ID() uuid.UUID {
	return u.id
}

func (u User) Email() string {
	return u.email
}

func (u User) Password() string {
	return u.password
}

func (u User) Firstname() string {
	return u.firstname
}

func (u User) Lastname() string {
	return u.lastname
}

func (u User) DateOfBirth() time.Time {
	return u.dateOfBirth
}

func (u User) CreatedAt() time.Time {
	return u.createdAt
}

func (u User) UpdatedAt() time.Time {
	return u.updatedAt
}
