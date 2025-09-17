package httpserver

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"todolist/todo"

	"github.com/gorilla/mux"
)

type HTTPHandlers struct {
	todoList *todo.List
}

func NewHTTPhandler(todoList *todo.List) *HTTPHandlers {

	return &HTTPHandlers{
		todoList: todoList,
	}
}

func (h *HTTPHandlers) HandleCreateTask(w http.ResponseWriter, r *http.Request) {

	var taskDTO TaskDTO

	if err := json.NewDecoder(r.Body).Decode(&taskDTO); err != nil {
		if errors.Is(err, todo.ErrTaskNotFound) {
			writeError(w, http.StatusNotFound, err)
		} else {
			writeError(w, http.StatusInternalServerError, err)
		}
		return
	}
	if err := taskDTO.ValidationForCreate(); err != nil {
		if errors.Is(err, todo.ErrTaskNotFound) {
			writeError(w, http.StatusNotFound, err)
		} else {
			writeError(w, http.StatusInternalServerError, err)
		}
		return
	}

	todoTask := todo.NewTask(taskDTO.Title, taskDTO.Description)
	if err := h.todoList.AddTask(todoTask); err != nil {
		if errors.Is(err, todo.ErrTaskNotFound) {
			writeError(w, http.StatusNotFound, err)
		} else {
			writeError(w, http.StatusBadRequest, err)
		}
		return
	}

	b, err := json.MarshalIndent(taskDTO, "", "	")
	if err != nil {
		panic(err)

	}

	w.WriteHeader(http.StatusCreated)
	if _, err := w.Write(b); err != nil {
		fmt.Println("failed to write response :", err)
	}

}
func (h *HTTPHandlers) HandleGetTask(w http.ResponseWriter, r *http.Request) {
	title := mux.Vars(r)["title"]

	task, err := h.todoList.GetTask(title)
	if err != nil {
		if errors.Is(err, todo.ErrTaskNotFound) {
			writeError(w, http.StatusNotFound, err)
		} else {
			writeError(w, http.StatusInternalServerError, err)
		}
		return
	}

	b, err := json.MarshalIndent(task, "", "	")
	if err != nil {
		panic(err)
	}
	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(b); err != nil {
		fmt.Println("failed to write response :", err)
		return
	}

}

func (h *HTTPHandlers) HandleGetAllTasks(w http.ResponseWriter, r *http.Request) {
	tasks := h.todoList.ListTasks()
	b, err := json.MarshalIndent(tasks, "", "	")
	if err != nil {
		panic(err)
	}
	if _, err := w.Write(b); err != nil {
		fmt.Println("failed to write response :", err)
		return
	}
}
func (h *HTTPHandlers) HandleGetAllUncompleteTasks(w http.ResponseWriter, r *http.Request) {
	uncompletedtasks := h.todoList.ListUnCompletedTasks()
	b, err := json.MarshalIndent(uncompletedtasks, "", "	")
	if err != nil {
		panic(err)
	}
	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(b); err != nil {
		fmt.Println("failed to write response :", err)
		return
	}
}
func (h *HTTPHandlers) HandleCompleteTask(w http.ResponseWriter, r *http.Request) {

	var CompleteDTO CompletedTaskDTO

	if err := json.NewDecoder(r.Body).Decode(&CompleteDTO); err != nil {
		if errors.Is(err, todo.ErrTaskNotFound) {
			writeError(w, http.StatusNotFound, err)
		} else {
			writeError(w, http.StatusInternalServerError, err)
		}
		return
	}
	title := mux.Vars(r)["title"]

	var (
		changedTask todo.Task
		err         error
	)

	if CompleteDTO.Complete {
		changedTask, err = h.todoList.CompletedTask(title)
	} else {
		changedTask, err = h.todoList.UnCompleteTask(title)
	}

	if err != nil {
		if errors.Is(err, todo.ErrTaskNotFound) {
			writeError(w, http.StatusNotFound, err)
		} else {
			writeError(w, http.StatusInternalServerError, err)
		}
		return
	}

	b, err := json.MarshalIndent(changedTask, "", "	")
	if err != nil {
		panic(err)
	}
	if _, err := w.Write(b); err != nil {
		fmt.Println("failed to write http response ")
		return
	}

}

func (h *HTTPHandlers) HandleDeleteTask(w http.ResponseWriter, r *http.Request) {
	title := mux.Vars(r)["title"]

	if err := h.todoList.DeleteTask(title); err != nil {
		if errors.Is(err, todo.ErrTaskNotFound) {
			writeError(w, http.StatusNotFound, err)
		} else {
			writeError(w, http.StatusInternalServerError, err)
		}
		return

	}
	w.WriteHeader(http.StatusNoContent)
}
