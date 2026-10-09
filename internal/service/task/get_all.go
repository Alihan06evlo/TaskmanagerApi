package task

import (
	"taskmanagerApi/internal/entity"
)

// Метод для получения всех задач
func (s *TaskService) GetAll() ([]entity.Task, error) {
	// Вызываем метод репозитория, который идет в базу данных
	tasks, err := s.repo.GetAll()
	if err != nil {
		return nil, err
	}

	return tasks, nil
}