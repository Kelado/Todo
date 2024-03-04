package storage

import "todo/internal/models"

type Storage interface {
	Find(id int64) (*models.Instance, error)
	Create(content string) (*models.Instance, error)
	List() ([]models.Instance, error)
	Search(term string) ([]models.Instance, error)
	Delete(id int64) (*models.Instance, error)
	Update(todo *models.Instance) (*models.Instance, error)
	SetCompleted(id int64, completed bool) error
}
