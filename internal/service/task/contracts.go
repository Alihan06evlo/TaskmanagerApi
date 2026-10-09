package task

import (
	"context"
	"taskmanagerApi/internal/entity"
)
type Repo interface {
	Create(ctx context.Context,t *entity.Task) (*entity.Task, error)
	Delete(id int) error
	GetById(id int) (*entity.Task,error)
	GetAll() ([]entity.Task, error)
}