package storage

import "109HW/internal/scheduler"

type Storage interface {
    Save(task *scheduler.Task)
    NextID() int
}
