package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"path"
	"strconv"
	"strings"

	"github.com/AlekseyKaramyshev/tasks-api/internal/models"
	"github.com/AlekseyKaramyshev/tasks-api/internal/storage"
)

type Handler struct{ Store storage.Storage }

func New(s storage.Storage) *Handler { return &Handler{Store: s} }

type ErrorResponse struct {
	Error string `json:"error"`
}

func writeJSON(w http.ResponseWriter, status int, v interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(v)
}

func writeError(w http.ResponseWriter, status int, errMsg string) {
	writeJSON(w, status, ErrorResponse{Error: errMsg})
}

func parseID(r *http.Request) (int, error) {
	idStr := strings.TrimPrefix(path.Base(r.URL.Path), "tasks/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return 0, errors.New("invalid task id")
	}
	return id, nil
}

// /tasks
func (h *Handler) TasksCollection(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		tasks := h.Store.List()
		writeJSON(w, http.StatusOK, tasks)

	case http.MethodPost:
		var task models.Task
		if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
			writeError(w, http.StatusBadRequest, "invalid JSON")
			return
		}

		if task.Title == "" {
			writeError(w, http.StatusBadRequest, "title is required")
			return
		}

		created, err := h.Store.Create(task)
		if err != nil {
			writeError(w, http.StatusInternalServerError, err.Error())
			return
		}

		writeJSON(w, http.StatusCreated, created)

	default:
		writeError(w, http.StatusMethodNotAllowed, "method not allowed")
	}
}

// /tasks/{id}
func (h *Handler) TaskItem(w http.ResponseWriter, r *http.Request) {
	id, err := parseID(r)
	if err != nil {
		writeError(w, http.StatusBadRequest, err.Error())
		return
	}

	switch r.Method {
	case http.MethodGet:
		task, exists := h.Store.Get(id)
		if !exists {
			writeError(w, http.StatusNotFound, "task not found")
			return
		}
		writeJSON(w, http.StatusOK, task)

	case http.MethodPut:
		var task models.Task
		if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
			writeError(w, http.StatusBadRequest, "invalid JSON")
			return
		}

		if task.Title == "" {
			writeError(w, http.StatusBadRequest, "title is required")
			return
		}

		updated, err := h.Store.Update(id, task)
		if err != nil {
			if errors.Is(err, storage.ErrNotFound) {
				writeError(w, http.StatusNotFound, "task not found")
			} else {
				writeError(w, http.StatusInternalServerError, err.Error())
			}
			return
		}
		writeJSON(w, http.StatusOK, updated)

	case http.MethodDelete:
		if err := h.Store.Delete(id); err != nil {
			if errors.Is(err, storage.ErrNotFound) {
				writeError(w, http.StatusNotFound, "task not found")
			} else {
				writeError(w, http.StatusInternalServerError, err.Error())
			}
			return
		}
		w.WriteHeader(http.StatusNoContent)

	default:
		writeError(w, http.StatusMethodNotAllowed, "method not allowed")
	}
}
