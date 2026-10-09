package task

import (
	"context"
	"taskmanagerApi/internal/entity"
	"taskmanagerApi/internal/store"

	"github.com/sirupsen/logrus"
)

func (s *TaskService) Create(ctx context.Context,t *entity.TaskInput) (*entity.Task, error){
	task := purchase(*t)
	return s.repo.Create(ctx,&task)
}

func purchase(t entity.TaskInput) entity.Task{
	if t.Status == "" {
		t.Status = "New"
	}
	if t.Score < 0 || t.Score > 100 {
		logrus.WithError(store.ErrValidation).Warn("[service] Invalid JSON")
		return entity.Task{}

	}
	return entity.Task{Title: t.Title, Status: t.Status,Score: t.Score}
}

