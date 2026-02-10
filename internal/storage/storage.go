package storage

import (
	"errors"
	"github.com/AlekseyKaramyshev/tasks-api/internal/models"
)

var ErrNotFound = errors.New("task not found")

type Storage interface {
	List() []models.Task
	Create(models.Task) (models.Task, error)
	Get(id int) (models.Task, bool)
	Update(id int, m models.Task) (models.Task, error)
	Delete(id int) error
}
