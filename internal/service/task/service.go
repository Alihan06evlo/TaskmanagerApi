package task

type TaskService struct{ repo Repo }

func NewTaskService(repo Repo) *TaskService { return &TaskService{repo: repo} }