package scheduler

import (
    "time"
)

type Task struct {
    ID           int
    Name         string
    ScheduleTime time.Time
    Command      string
    Status       string
}
