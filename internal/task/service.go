package task

import (
	"errors"
	"strings"
)

var (
	ErrTaskNotFound  = errors.New("task not found")
	ErrTitleRequired = errors.New("title is required")
)

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) ListTasks() []Task {
	return s.repo.List()
}

func (s *Service) GetTask(id int) (Task, error) {
	task, ok := s.repo.GetByID(id)
	if !ok {
		return Task{}, ErrTaskNotFound
	}

	return task, nil
}

func (s *Service) CreateTask(input CreateTaskInput) (Task, error) {
	normalizedInput, err := normalizeCreateInput(input)
	if err != nil {
		return Task{}, err
	}

	return s.repo.Create(normalizedInput), nil
}

func (s *Service) UpdateTask(id int, input UpdateTaskInput) (Task, error) {
	normalizedInput, err := normalizeUpdateInput(input)
	if err != nil {
		return Task{}, err
	}

	task, ok := s.repo.Update(id, normalizedInput)
	if !ok {
		return Task{}, ErrTaskNotFound
	}

	return task, nil
}

func (s *Service) DeleteTask(id int) error {
	if ok := s.repo.Delete(id); !ok {
		return ErrTaskNotFound
	}

	return nil
}

func normalizeCreateInput(input CreateTaskInput) (CreateTaskInput, error) {
	title := strings.TrimSpace(input.Title)
	if title == "" {
		return CreateTaskInput{}, ErrTitleRequired
	}

	return CreateTaskInput{
		Title:       title,
		Description: strings.TrimSpace(input.Description),
		Done:        input.Done,
	}, nil
}

func normalizeUpdateInput(input UpdateTaskInput) (UpdateTaskInput, error) {
	title := strings.TrimSpace(input.Title)
	if title == "" {
		return UpdateTaskInput{}, ErrTitleRequired
	}

	return UpdateTaskInput{
		Title:       title,
		Description: strings.TrimSpace(input.Description),
		Done:        input.Done,
	}, nil
}
