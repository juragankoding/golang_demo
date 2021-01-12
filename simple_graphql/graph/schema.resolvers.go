package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/juragankoding/simple_graphql/graph/generated"
	"github.com/juragankoding/simple_graphql/graph/model"
)

func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	todoInput := &model.Todo{
		Text: input.Text,
	}

	return r.ServiceTodo.AddTodo(todoInput)
}

func (r *mutationResolver) UdpateStatusTodo(ctx context.Context, id string) (*model.Message, error) {
	todo, err := r.ServiceTodo.GetTodo(id)

	if err != nil {
		return nil, err
	}

	todo.Done = true

	err = r.ServiceTodo.UpdateTodo(todo)

	if err != nil {
		return nil, err
	}

	return &model.Message{
		Message: "Success",
	}, nil
}

func (r *mutationResolver) UpdateTodo(ctx context.Context, input model.UpdateTodo) (*model.Todo, error) {
	todoInput := &model.Todo{
		ID:   input.ID,
		Text: input.Text,
		Done: input.Done,
	}
	err := r.ServiceTodo.UpdateTodo(todoInput)

	if err != nil {
		return nil, err
	} else {
		return todoInput, nil
	}
}

func (r *mutationResolver) DeleteTodo(ctx context.Context, id string) (*model.Message, error) {
	if err := r.ServiceTodo.DeleteTodo(id); err != nil {
		return nil, err
	} else {
		return &model.Message{
			Message: "Delete success",
		}, nil
	}
}

func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	return r.ServiceTodo.ListTodo()
}

func (r *queryResolver) Status(ctx context.Context) (*model.Message, error) {
	return &model.Message{
		Message: "this is demo of graphql on golang",
	}, nil
}

func (r *queryResolver) SingleTodo(ctx context.Context, id string) (*model.Todo, error) {
	return r.ServiceTodo.GetTodo(id)
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
