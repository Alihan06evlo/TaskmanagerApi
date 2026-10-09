package entity

type Task struct {
	ID int `json:"id"`
	Title string `json:"title"`
	Status string `json:"status"`
	Score int `json:"score"`
}

type TaskInput struct {
	Title string `json:"title"`
	Status string `json:"status"`
	Score int `json:"score"`
}