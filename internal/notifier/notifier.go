package notifier

import (
	sch "109HW/internal/scheduler"
	"log"
)

type Notifier interface {
	Notify(task *sch.Task, message string)
}

type EmailNotifier struct{}

func NewNotifier() *EmailNotifier {
	return &EmailNotifier{}
}

func (e *EmailNotifier) Notify(task *sch.Task, message string) {
	log.Printf("Notification for task %d: %s", task.ID, message)
}
