package repository

import (
	"golang-echo-graphql-example/graph/model"
)

type messageRepositoryImpl struct {
	messages []*model.Message
}

func NewMessageInMemoryRepository() MessageRepository {
	return &messageRepositoryImpl{
		messages: make([]*model.Message, 0),
	}
}

func (messageRepo *messageRepositoryImpl) Save(message *model.Message) (*model.Message, error) {
	messageRepo.messages = append(messageRepo.messages, message)
	return message, nil
}
