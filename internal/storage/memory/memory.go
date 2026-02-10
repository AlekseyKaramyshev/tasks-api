package memory

import (
	"github.com/AlekseyKaramyshev/tasks-api/internal/models"
	"github.com/AlekseyKaramyshev/tasks-api/internal/storage"
	"sync"
)

type Storage struct {
	mu     sync.RWMutex
	tasks  map[int]models.Task
	nextID int
}

func New() *Storage {
	return &Storage{
		tasks:  make(map[int]models.Task),
		nextID: 1,
	}
}

func (s *Storage) List() []models.Task {
	s.mu.RLock()
	defer s.mu.RUnlock()

	tasks := make([]models.Task, 0, len(s.tasks))
	for _, task := range s.tasks {
		tasks = append(tasks, task)
	}
	return tasks
}

func (s *Storage) Create(task models.Task) (models.Task, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	task.ID = s.nextID
	task.SetDefaults()
	s.nextID++

	s.tasks[task.ID] = task
	return task, nil
}

func (s *Storage) Get(id int) (models.Task, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	task, exists := s.tasks[id]
	return task, exists
}

func (s *Storage) Update(id int, task models.Task) (models.Task, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, exists := s.tasks[id]; !exists {
		return models.Task{}, storage.ErrNotFound
	}

	task.ID = id
	s.tasks[id] = task
	return task, nil
}

func (s *Storage) Delete(id int) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, exists := s.tasks[id]; !exists {
		return storage.ErrNotFound
	}

	delete(s.tasks, id)
	return nil
}
