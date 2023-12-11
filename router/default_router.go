package router

import (
	"encoding/json"
	"net/http"
	"todolist/dto"
)

func HeaderConfig(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
}

func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	HeaderConfig(w)

	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(dto.TodoResponse{
		Success: false,
		Message: "service not found",
		Data:    nil,
	})
}
