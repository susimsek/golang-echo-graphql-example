package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"golang-echo-graphql-example/config"
	"golang-echo-graphql-example/graph"
	"golang-echo-graphql-example/repository"
	"golang-echo-graphql-example/routes"
	"golang-echo-graphql-example/service"
)

var userService service.UserService
var messageService service.MessageService

func main() {
	e := echo.New()
	config.CORSConfig(e)

	r := graph.NewResolver(userService, messageService)
	srv := config.NewGraphQLServer(r)

	routes.GetActuatorRoutes(e)
	routes.GetGraphqlRoutes(e, srv)

	// echo server 9000 de başlatıldı.
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", config.ServerPort)))
}

func init() {
	userRepository := repository.NewUserInMemoryRepository()
	userService = service.NewUserService(userRepository)
	messageRepository := repository.NewMessageInMemoryRepository()
	messageService = service.NewMessageService(messageRepository)
}
