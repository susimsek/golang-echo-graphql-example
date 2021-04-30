package service

import (
	uuid "github.com/satori/go.uuid"
	"golang-echo-graphql-example/graph/model"
	"golang-echo-graphql-example/repository"
	"time"
)

type MessageService interface {
	Save(message *model.Message) (*model.Message, error)
}

type messageServiceImpl struct {
	messageRepo repository.MessageRepository
}

func NewMessageService(messageRepo repository.MessageRepository) MessageService {
	return &messageServiceImpl{messageRepo: messageRepo}
}

func (messageService *messageServiceImpl) Save(message *model.Message) (*model.Message, error) {
	id := uuid.NewV4().String()
	message.CreatedAt = time.Now()
	message.ID = id
	return messageService.messageRepo.Save(message)
}
