package main

import (
    "109HW/internal/scheduler"
    "109HW/internal/storage"
    "109HW/internal/notifier"
    "log"
)

func main() {
    taskStorage := storage.NewInMemoryStorage()

    taskNotifier := notifier.NewNotifier()

    taskScheduler := scheduler.NewScheduler(taskStorage, taskNotifier)

    err := taskScheduler.AddTask("Example Task", "*/5 * * * *", "echo 'Hello, World!'")
    if err != nil {
        log.Fatalf("Error scheduling task: %v", err)
    }

    taskScheduler.Start()
}
