package todos

import "todo/internal/models"

type Service interface {
	Find(id int64) (models.Instance, error)
	Create(todo *models.Instance) error
	List() ([]models.Instance, error)
	Search(term string) ([]models.Instance, error)
	Delete(id int64) error
	Update(todo *models.Instance) error
	SetCompleted(id int64, completed bool) error
}
