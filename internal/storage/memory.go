package storage

import "109HW/internal/scheduler"

type InMemoryStorage struct {
    tasks map[int]*scheduler.Task
    nextID int
}

func NewInMemoryStorage() *InMemoryStorage {
    return &InMemoryStorage{
        tasks:  make(map[int]*scheduler.Task),
        nextID: 1,
    }
}

func (s *InMemoryStorage) Save(task *scheduler.Task) {
    s.tasks[task.ID] = task
}

func (s *InMemoryStorage) NextID() int {
    id := s.nextID
    s.nextID++
    return id
}
