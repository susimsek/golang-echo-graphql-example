package repository

import (
	"golang-echo-graphql-example/graph/model"
)

type MessageRepository interface {
	Save(message *model.Message) (*model.Message, error)
}
