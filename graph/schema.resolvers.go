package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"golang-echo-graphql-example/graph/generated"
	"golang-echo-graphql-example/graph/model"
)

func (r *mutationResolver) CreateUser(ctx context.Context, input model.UserInput) (*model.User, error) {
	user := &model.User{
		FirstName: input.FirstName,
		LastName:  input.LastName,
		Email:     input.Email,
	}

	createdUser, err := r.userService.Save(user)
	if err != nil {
		return nil, err
	}
	return createdUser, nil
}

func (r *mutationResolver) UpdateUser(ctx context.Context, id string, input model.UserInput) (*model.User, error) {
	user := &model.User{
		FirstName: input.FirstName,
		LastName:  input.LastName,
		Email:     input.Email,
	}
	user, err := r.userService.Update(id, user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *mutationResolver) DeleteUser(ctx context.Context, id string) (bool, error) {
	err := r.userService.DeleteById(id)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (r *mutationResolver) CreateMessage(ctx context.Context, input model.MessageInput) (*model.Message, error) {
	message := &model.Message{
		Message: input.Message,
	}
	createdMessage, err := r.messageService.Save(message)
	if err != nil {
		return nil, err
	}

	r.mu.Lock()
	r.messageChan <- createdMessage
	r.mu.Unlock()
	return createdMessage, nil
}

func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	return r.userService.FindAll()
}

func (r *queryResolver) User(ctx context.Context, id string) (*model.User, error) {
	user, err := r.userService.FindById(id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *subscriptionResolver) MessageAdded(ctx context.Context) (<-chan *model.Message, error) {
	return r.messageChan, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// Subscription returns generated.SubscriptionResolver implementation.
func (r *Resolver) Subscription() generated.SubscriptionResolver { return &subscriptionResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type subscriptionResolver struct{ *Resolver }
