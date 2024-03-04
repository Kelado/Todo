package storage

import (
	"errors"
	"strings"
	"todo/internal/models"
)

type DumbStorage struct {
	ID    int64
	store []models.Instance
}

func NewDumpStorage() *DumbStorage {
	store := []models.Instance{
		{1, "content 1", false},
		{2, "content 2", false},
		{3, "content 3", false},
	}

	return &DumbStorage{
		ID:    4,
		store: store,
	}
}

// func (s *DumbStorage) GetById(id int64) (todos.Instance, error) {
// 	for _, todo := range s.store {
// 		if todo.Id == id {
// 			return todo, nil
// 		}
// 	}
// 	return todos.Instance{}, errors.New("todo not found")
// }

func (s *DumbStorage) Find(id int64) (*models.Instance, error) {
	for i, t := range s.store {
		if t.Id == id {
			return &s.store[i], nil
		}
	}
	return nil, errors.New("could not find todo item")
}

func (s *DumbStorage) Create(content string) (*models.Instance, error) {
	t := models.Instance{
		Id:        int64(s.ID),
		Content:   content,
		Completed: false,
	}
	s.store = append(s.store, t)

	s.ID++

	return &t, nil
}

func (s *DumbStorage) List() ([]models.Instance, error) {
	return s.store, nil
}

func (s *DumbStorage) Search(term string) ([]models.Instance, error) {
	var list []models.Instance
	for _, todo := range s.store {
		if strings.Contains(todo.Content, term) {
			list = append(list, todo)
		}
	}
	return list, nil
}

func (s *DumbStorage) Delete(id int64) (*models.Instance, error) {
	var newTodos []models.Instance
	var todoToDelete models.Instance
	for _, todo := range s.store {
		if todo.Id == id {
			todoToDelete = todo
			continue
		}
		newTodos = append(newTodos, todo)
	}
	s.store = newTodos
	return &todoToDelete, nil
}

func (s *DumbStorage) Update(todo *models.Instance) (*models.Instance, error) {
	var updatedTodo models.Instance
	for i, t := range s.store {
		if t.Id == todo.Id {
			s.store[i].Content = todo.Content
			updatedTodo = s.store[i]
			break
		}
	}
	return &updatedTodo, nil
}

func (s *DumbStorage) SetCompleted(id int64, completed bool) error {
	t, _ := s.Find(id)
	t.Completed = true
	return nil
}
