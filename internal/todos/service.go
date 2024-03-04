package todos

import (
	"fmt"

	"todo/internal/models"
	"todo/internal/storage"
)

var ID = 5

type service struct {
	db storage.Storage
}

func NewService(db storage.Storage) *service {
	return &service{
		db: db,
	}
}

func (s *service) Find(id int64) (*models.Instance, error) {
	todo, err := s.db.Find(id)
	if err != nil {
		fmt.Println(err)
	}
	return todo, err
}

func (s *service) Create(content string) (*models.Instance, error) {
	t, err := s.db.Create(content)
	return t, err
}

func (s *service) List() ([]models.Instance, error) {
	l, err := s.db.List()
	return l, err
}

func (s *service) Search(term string) ([]models.Instance, error) {
	l, err := s.db.Search(term)
	return l, err
}

func (s *service) Delete(id int64) (*models.Instance, error) {
	t, err := s.db.Delete(id)
	return t, err
}

func (s *service) Update(todo *models.Instance) (*models.Instance, error) {
	t, err := s.db.Update(todo)
	return t, err
}

func (s *service) SetCompleted(id int64, completed bool) error {
	err := s.db.SetCompleted(id, completed)
	return err
}
