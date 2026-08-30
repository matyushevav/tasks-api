package storage

import (
	"fmt"
	"sync"
	"tasks-api/internal/models"
	"time"
)

// MemoryStorage - хранилище задач в оперативной памяти
type MemoryStorage struct {
	tasks  map[int]models.Task // мапа для хранения задач (ключ - ID)
	nextID int                 // счетчик для генерации новых ID
	mu     sync.RWMutex        // мьютекс для потокобезопасности
}

// NewMemoryStorage - конструктор, создает новое хранилище
func NewMemoryStorage() *MemoryStorage {
	return &MemoryStorage{
		tasks:  make(map[int]models.Task),
		nextID: 1, // начинаем ID с 1
	}
}

// List - возвращает список всех задач
func (s *MemoryStorage) List() []models.Task {
	s.mu.RLock()
	defer s.mu.RUnlock()

	result := make([]models.Task, 0, len(s.tasks))
	for _, task := range s.tasks {
		result = append(result, task)
	}
	return result
}

// Create - создает новую задачу
func (s *MemoryStorage) Create(task models.Task) (models.Task, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	task.ID = s.nextID
	s.nextID++

	// Генерируем время создания
	task.CreatedAt = time.Now().Format(time.RFC3339)

	s.tasks[task.ID] = task
	return task, nil
}

// Get - возвращает задачу по ID
func (s *MemoryStorage) Get(id int) (models.Task, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	task, exists := s.tasks[id]
	return task, exists
}

// Update - обновляет задачу
func (s *MemoryStorage) Update(id int, task models.Task) (models.Task, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	existing, exists := s.tasks[id]
	if !exists {
		return models.Task{}, fmt.Errorf("task with id %d not found", id)
	}

	// Обновляем только title и done, created_at оставляем старый
	existing.Title = task.Title
	existing.Done = task.Done

	s.tasks[id] = existing
	return existing, nil
}

// Delete - удаляет задачу по ID
func (s *MemoryStorage) Delete(id int) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, exists := s.tasks[id]; !exists {
		return fmt.Errorf("task with id %d not found", id)
	}

	delete(s.tasks, id)
	return nil
}
