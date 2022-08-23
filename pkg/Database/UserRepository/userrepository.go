package UserRepository

import (
	"main.go/model"
	"main.go/pkg/Database"
)

type UserRepository struct {
	store *Database.Store
}

func (ur *UserRepository) Create(u *model.User) (*model.User, error) {
	return &model.User{
		ID:       0,
		Email:    "",
		Username: "",
		Password: "",
	}, nil
}
