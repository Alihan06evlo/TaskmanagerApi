package task

import "taskmanagerApi/internal/store"

type Repository struct{ store *store.Store }

func NewRepository(store *store.Store) *Repository { return &Repository{store: store} }

