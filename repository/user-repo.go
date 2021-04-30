package repository

import (
	"golang-echo-graphql-example/graph/model"
)

type UserRepository interface {
	Count() int64
	FindAll() ([]*model.User, error)
	Save(user *model.User) (*model.User, error)
	FindById(id string) (*model.User, error)
	Update(id string, user *model.User) (*model.User, error)
	DeleteById(id string) error
}
