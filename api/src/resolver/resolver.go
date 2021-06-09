package resolver

import (
	"api/graph/model"
	"strconv"
	"sync"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.`

type todoStore struct {
	Data      []*model.Todo
	IdCounter int
	mutex     sync.Mutex
}

var store *todoStore

func init() {
	initialTodos := []*model.Todo{
		{
			ID:   "1",
			Text: "text",
			Done: false,
			User: &model.User{
				ID:   "1",
				Name: "aoba",
			},
		},
		{
			ID:   "2",
			Text: "text2",
			Done: true,
			User: &model.User{
				ID:   "2",
				Name: "aoba2",
			},
		},
	}

	store = &todoStore{
		Data:      initialTodos,
		IdCounter: 3,
	}
}

func (s *todoStore) AddTodo(t *model.Todo) *model.Todo {
	s.mutex.Lock()
	t.ID = strconv.Itoa(s.IdCounter)
	s.Data = append(s.Data, t)
	s.IdCounter = s.IdCounter + 1
	s.mutex.Unlock()

	return t
}

type Resolver struct {
	TodoAddedChans map[string]chan *model.Todo
}

func New() *Resolver {
	return &Resolver{TodoAddedChans: map[string]chan *model.Todo{}}
}
