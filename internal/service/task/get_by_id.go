package task

import "taskmanagerApi/internal/entity"

func (s *TaskService) GetById(id int) (*entity.Task, error) {
	return s.repo.GetById(id)
}