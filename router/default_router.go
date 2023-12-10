package router

import (
	"encoding/json"
	"net/http"
	"todolist/dto"
)

func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(dto.TodoResponse{
		Success: false,
		Message: "service not found",
		Data:    nil,
	})
}
