package service

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/juragankoding/simple_graphql/graph/model"
)

type TodoServices struct {
	todos []*model.Todo
}

func (ts *TodoServices) AddTodo(todo *model.Todo) (*model.Todo, error) {
	var lastid = 0
	var err error

	if len(ts.todos) > 0 {
		todoLast := ts.todos[len(ts.todos)-1]

		lastid, err = strconv.Atoi(todoLast.ID)

		if err != nil {
			return nil, err
		}
	}

	todo.ID = strconv.Itoa(lastid + 1)

	ts.todos = append(ts.todos, todo)

	return todo, nil
}

func (ts *TodoServices) GetTodo(id string) (*model.Todo, error) {
	for todo := range ts.todos {
		if *&ts.todos[todo].ID == id {
			return *&ts.todos[todo], nil
		}
	}

	return nil, errors.New("not have data todo in todos variable")
}

func (ts *TodoServices) DeleteTodo(id string) error {
	for todo := range ts.todos {
		if *&ts.todos[todo].ID == id {
			fmt.Print(todo)

			ts.todos = append(ts.todos[:todo], ts.todos[todo+1:]...)

			return nil
		}
	}

	return errors.New("not have data todo to delete in todos variable")
}

func (ts *TodoServices) UpdateTodo(todoInput *model.Todo) error {
	for todo := range ts.todos {
		if *&ts.todos[todo].ID == *&todoInput.ID {
			*&ts.todos[todo] = todoInput
			return nil
		}
	}

	return errors.New("not have data todo to update in todos variable")
}

func (ts *TodoServices) ListTodo() ([]*model.Todo, error) {
	return ts.todos, nil
}
