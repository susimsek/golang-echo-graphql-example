package service

import (
	uuid "github.com/satori/go.uuid"
	"golang-echo-graphql-example/graph/model"
	"golang-echo-graphql-example/repository"
	"time"
)

type UserService interface {
	FindAll() ([]*model.User, error)
	Save(user *model.User) (*model.User, error)
	FindById(id string) (*model.User, error)
	Update(id string, user *model.User) (*model.User, error)
	DeleteById(id string) error
}

type userServiceImpl struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return &userServiceImpl{userRepo: userRepo}
}

func (userService *userServiceImpl) FindAll() ([]*model.User, error) {
	return userService.userRepo.FindAll()
}

func (userService *userServiceImpl) Save(user *model.User) (*model.User, error) {
	id := uuid.NewV4().String()
	user.CreatedAt = time.Now()
	user.ID = id
	return userService.userRepo.Save(user)
}

func (userService *userServiceImpl) FindById(id string) (*model.User, error) {
	return userService.userRepo.FindById(id)
}

func (userService *userServiceImpl) Update(id string, user *model.User) (*model.User, error) {
	return userService.userRepo.Update(id, user)
}

func (userService *userServiceImpl) DeleteById(id string) error {
	return userService.userRepo.DeleteById(id)
}
