package repository

import (
	"context"
	"golang-echo-graphql-example/exception"
	"golang-echo-graphql-example/graph/model"
)

var ctx context.Context = context.Background()

type userRepositoryImpl struct {
	users []*model.User
}

func NewUserInMemoryRepository() UserRepository {
	return &userRepositoryImpl{
		users: make([]*model.User, 0),
	}
}

func (userRepo *userRepositoryImpl) Count() int64 {
	return int64(len(userRepo.users))
}

func (userRepo *userRepositoryImpl) FindAll() ([]*model.User, error) {
	return userRepo.users, nil
}

func (userRepo *userRepositoryImpl) Save(user *model.User) (*model.User, error) {
	userRepo.users = append(userRepo.users, user)
	return user, nil
}

func (userRepo *userRepositoryImpl) FindById(id string) (*model.User, error) {
	i, found := searchUserById(userRepo.users, id)
	if !found {
		return nil, exception.ResourceNotFoundException("User", "id", id)
	}
	existUser := userRepo.users[i]
	return existUser, nil
}

func (userRepo *userRepositoryImpl) Update(id string, user *model.User) (*model.User, error) {
	i, found := searchUserById(userRepo.users, id)
	if !found {
		return nil, exception.ResourceNotFoundException("User", "id", id)
	}
	existUser := userRepo.users[i]
	existUser.FirstName = user.FirstName
	existUser.LastName = user.LastName
	existUser.Email = user.Email
	userRepo.users[i] = existUser

	return existUser, nil
}

func (userRepo *userRepositoryImpl) DeleteById(id string) error {
	i, found := searchUserById(userRepo.users, id)
	if !found {
		return exception.ResourceNotFoundException("User", "id", id)
	}
	userRepo.users = append(userRepo.users[:i], userRepo.users[i+1:]...)
	return nil
}

func searchUserById(users []*model.User, id string) (int, bool) {
	for i, v := range users {
		if v.ID == id {
			return i, true
		}
	}
	return -1, false
}
