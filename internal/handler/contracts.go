package task

import (
	"context"
	"taskmanagerApi/internal/entity"
)

type Service interface {
	Create(ctx context.Context,t *entity.TaskInput) (*entity.Task, error)
	Delete(id int) error
	GetById(id int) (*entity.Task,error)
	GetAll() ([]entity.Task, error)
}