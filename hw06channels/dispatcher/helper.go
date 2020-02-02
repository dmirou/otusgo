package dispatcher

import (
	"errors"
	"math/rand"
	"time"
)

// generateTasks makes a slice with success and error tasks.
func generateTasks(successCount, errorCount int) []Task {
	successTask := func() error {
		sleepTime := time.Duration(rand.Intn(500)) * time.Millisecond
		time.Sleep(sleepTime)
		return nil
	}
	errorTask := func() error {
		sleepTime := time.Duration(rand.Intn(500)) * time.Millisecond
		time.Sleep(sleepTime)
		return errors.New("there is a task error")
	}
	tasks := make([]Task, successCount+errorCount)
	if successCount+errorCount == 0 {
		return tasks
	}

	for i := 0; i < successCount; i++ {
		tasks[i] = successTask
	}
	for i := successCount; i < successCount+errorCount; i++ {
		tasks[i] = errorTask
	}
	rand.Shuffle(len(tasks), func(i, j int) {
		tasks[i], tasks[j] = tasks[j], tasks[i]
	})
	return tasks
}
