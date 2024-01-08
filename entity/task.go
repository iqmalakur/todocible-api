package entity

import "time"

type Task struct {
	Id        string    `json:"id"`
	Name      string    `json:"name"`
	DueDate   time.Time `json:"due_date"`
	Category  string    `json:"category"`
	Completed bool      `json:"completed"`
}
