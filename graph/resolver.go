package graph

import (
	"golang-echo-graphql-example/graph/model"
	"golang-echo-graphql-example/service"
	"sync"
)

type Resolver struct {
	userService    service.UserService
	messageService service.MessageService
	messageChan    chan *model.Message
	mu             sync.Mutex
}

func NewResolver(userService service.UserService, messageService service.MessageService) *Resolver {
	return &Resolver{
		userService:    userService,
		messageService: messageService,
		messageChan:    make(chan *model.Message),
	}
}
