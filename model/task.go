package model

type Task struct {
	ID          int    `json:"id" gorm:"primaryKey"`
	Title       string `json:"title"`
	Completed	bool	`json:"completed"`
	UserID   	int    `json:"user_id"`
}