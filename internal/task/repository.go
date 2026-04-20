package task

import (
	"sort"
	"sync"
	"time"
)

type Repository interface {
	List() []Task
	GetByID(id int) (Task, bool)
	Create(input CreateTaskInput) Task
	Update(id int, input UpdateTaskInput) (Task, bool)
	Delete(id int) bool
}

type MemoryRepository struct {
	mu     sync.Mutex
	tasks  map[int]Task
	nextID int
}

func NewMemoryRepository() *MemoryRepository {
	return &MemoryRepository{
		tasks:  make(map[int]Task),
		nextID: 1,
	}
}

func (r *MemoryRepository) List() []Task {
	r.mu.Lock()
	defer r.mu.Unlock()

	tasks := make([]Task, 0, len(r.tasks))
	for _, task := range r.tasks {
		tasks = append(tasks, task)
	}

	sort.Slice(tasks, func(i, j int) bool {
		return tasks[i].ID < tasks[j].ID
	})

	return tasks
}

func (r *MemoryRepository) GetByID(id int) (Task, bool) {
	r.mu.Lock()
	defer r.mu.Unlock()

	task, ok := r.tasks[id]
	return task, ok
}

func (r *MemoryRepository) Create(input CreateTaskInput) Task {
	r.mu.Lock()
	defer r.mu.Unlock()

	now := time.Now()
	task := Task{
		ID:          r.nextID,
		Title:       input.Title,
		Description: input.Description,
		Done:        input.Done,
		CreatedAt:   now,
		UpdatedAt:   now,
	}

	r.tasks[task.ID] = task
	r.nextID++

	return task
}

func (r *MemoryRepository) Update(id int, input UpdateTaskInput) (Task, bool) {
	r.mu.Lock()
	defer r.mu.Unlock()

	task, ok := r.tasks[id]
	if !ok {
		return Task{}, false
	}

	task.Title = input.Title
	task.Description = input.Description
	task.Done = input.Done
	task.UpdatedAt = time.Now()
	r.tasks[id] = task

	return task, true
}

func (r *MemoryRepository) Delete(id int) bool {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, ok := r.tasks[id]; !ok {
		return false
	}

	delete(r.tasks, id)
	return true
}
