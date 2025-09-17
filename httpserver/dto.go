package httpserver

import (
	"encoding/json"
	"errors"  
	"net/http"
	"time"
)

type CompletedTaskDTO struct {
	Complete bool
}

type TaskDTO struct {
	Title       string
	Description string
}

func (t *TaskDTO) ValidationForCreate() error {
	if t.Title == "" {
		return errors.New("title is empty")
	} else if t.Description == "" {
		return errors.New("description is empty")
	}
	return nil
}

type ErrorDTO struct {
	Masseg string    `json:"message"`
	Time   time.Time `json:"time"`
}

func writeError(w http.ResponseWriter, status int, err error) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	_ = json.NewEncoder(w).Encode(ErrorDTO{
		Masseg: err.Error(),
		Time:   time.Now(),
	})
}
