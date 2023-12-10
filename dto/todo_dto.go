package dto

type (
	TodoRequest struct {
		Title       string `json:"title"`
		Description string `json:"description"`
	}

	TodoResponse struct {
		Success bool   `json:"success"`
		Message string `json:"message"`
		Data    any    `json:"data"`
	}
)
