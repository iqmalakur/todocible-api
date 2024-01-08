package entity

import "time"

type Task struct {
	Id        string    `json:"id"`
	User      string    `json:"user"`
	Category  string    `json:"category"`
	Name      string    `json:"name"`
	DueDate   time.Time `json:"due_date"`
	Completed bool      `json:"completed"`
}
