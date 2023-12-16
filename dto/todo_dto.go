package dto

import "time"

type (
	TodoRequest struct {
		Title       string    `json:"title"`
		Description string    `json:"description"`
		DueDate     time.Time `json:"due_date"`
	}

	TodoResponse struct {
		Success bool   `json:"success"`
		Message string `json:"message"`
		Data    any    `json:"data"`
	}
)
