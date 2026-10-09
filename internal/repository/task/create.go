package task

import (
	"context"
	"taskmanagerApi/internal/entity"
	 "taskmanagerApi/internal/store"
)

func (r *Repository) Create(ctx context.Context,t *entity.Task) (*entity.Task, error){
	err := r.store.GetConn().QueryRowContext(ctx,`INSERT INTO tasks(title,status,score) VALUES ($1,$2,$3) RETURNING id,title,status,score`,t.Title,t.Status,t.Score).Scan(&t.ID,&t.Title,&t.Status,&t.Score)
	if err != nil {
		err,_ =store.DBErrToErr(err)
		return nil, err
	}
return t, nil
}