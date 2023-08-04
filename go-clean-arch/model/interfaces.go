package model

import (
	"github.com/google/uuid"
	"golang.org/x/net/context"
)

type UserService interface {
	Get(ctx context.Context, uid uuid.UUID) (*User, error)
	Signup(ctx context.Context, u *User) error
}

type UserRepository interface {
	FindById(ctx context.Context, uid uuid.UUID) (*User, error)
}
