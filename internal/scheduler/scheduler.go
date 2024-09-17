package scheduler

import (
    "109HW/internal/storage"
    "109HW/internal/notifier"
    "log"
    "time"
    "github.com/robfig/cron/v3"
)

type Scheduler struct {
    storage   storage.Storage
    notifier  notifier.Notifier
    cron      *cron.Cron
}

func NewScheduler(storage storage.Storage, notifier notifier.Notifier) *Scheduler {
    return &Scheduler{
        storage:  storage,
        notifier: notifier,
        cron:     cron.New(),
    }
}

func (s *Scheduler) AddTask(name, cronExpr, command string) error {
    id := s.storage.NextID()
    task := &Task{
        ID:           id,
        Name:         name,
        Command:      command,
        ScheduleTime: time.Now(),
        Status:       "Scheduled",
    }
    s.storage.Save(task)

    _, err := s.cron.AddFunc(cronExpr, func() {
        err := s.ExecuteTask(task)
        if err != nil {
            log.Printf("Task %d failed: %v", task.ID, err)
            task.Status = "Failed"
            s.notifier.Notify(task, "Task failed")
        } else {
            task.Status = "Completed"
            s.notifier.Notify(task, "Task completed")
        }
    })

    if err != nil {
        return err
    }

    return nil
}

func (s *Scheduler) ExecuteTask(task *Task) error {
    log.Printf("Executing task %d: %s", task.ID, task.Command)
    return nil
}

func (s *Scheduler) Start() {
    s.cron.Start()
}
