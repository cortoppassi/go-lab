package task

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"strings"

	"go-lab/internal/httpapi"
)

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/tasks", h.handleTasks)
	mux.HandleFunc("/tasks/", h.handleTaskByID)
}

func (h *Handler) handleTasks(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		httpapi.WriteJSON(w, http.StatusOK, h.service.ListTasks())
	case http.MethodPost:
		var input CreateTaskInput
		if err := decodeJSON(r, &input); err != nil {
			httpapi.WriteError(w, http.StatusBadRequest, err.Error())
			return
		}

		task, err := h.service.CreateTask(input)
		if err != nil {
			h.writeServiceError(w, err)
			return
		}

		httpapi.WriteJSON(w, http.StatusCreated, task)
	default:
		httpapi.WriteError(w, http.StatusMethodNotAllowed, "metodo nao permitido")
	}
}

func (h *Handler) handleTaskByID(w http.ResponseWriter, r *http.Request) {
	id, err := parseTaskID(r.URL.Path)
	if err != nil {
		httpapi.WriteError(w, http.StatusBadRequest, "id invalido")
		return
	}

	switch r.Method {
	case http.MethodGet:
		task, err := h.service.GetTask(id)
		if err != nil {
			h.writeServiceError(w, err)
			return
		}

		httpapi.WriteJSON(w, http.StatusOK, task)
	case http.MethodPut:
		var input UpdateTaskInput
		if err := decodeJSON(r, &input); err != nil {
			httpapi.WriteError(w, http.StatusBadRequest, err.Error())
			return
		}

		task, err := h.service.UpdateTask(id, input)
		if err != nil {
			h.writeServiceError(w, err)
			return
		}

		httpapi.WriteJSON(w, http.StatusOK, task)
	case http.MethodDelete:
		if err := h.service.DeleteTask(id); err != nil {
			h.writeServiceError(w, err)
			return
		}

		httpapi.WriteJSON(w, http.StatusOK, map[string]string{
			"message": "tarefa removida com sucesso",
		})
	default:
		httpapi.WriteError(w, http.StatusMethodNotAllowed, "metodo nao permitido")
	}
}

func (h *Handler) writeServiceError(w http.ResponseWriter, err error) {
	switch {
	case errors.Is(err, ErrTitleRequired):
		httpapi.WriteError(w, http.StatusBadRequest, "o campo title e obrigatorio")
	case errors.Is(err, ErrTaskNotFound):
		httpapi.WriteError(w, http.StatusNotFound, "tarefa nao encontrada")
	default:
		httpapi.WriteError(w, http.StatusInternalServerError, "erro interno")
	}
}

func decodeJSON(r *http.Request, target any) error {
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	if err := decoder.Decode(target); err != nil {
		return errors.New("json invalido")
	}

	return nil
}

func parseTaskID(path string) (int, error) {
	idText := strings.TrimPrefix(path, "/tasks/")
	if idText == "" || strings.Contains(idText, "/") {
		return 0, errors.New("invalid id")
	}

	return strconv.Atoi(idText)
}
