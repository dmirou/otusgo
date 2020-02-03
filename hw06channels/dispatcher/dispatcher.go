package dispatcher

import "fmt"

type Task func() error

func Run(tasks []Task, N int, M int) error {

	tasksCount := len(tasks)
	tasksQueue := make(chan Task, N)
	quitChan := make(chan bool)
	resultsChan := make(chan error, N)

	go tasksToQueue(quitChan, tasks, tasksQueue)
	for i := 0; i < N; i++ {
		go worker(quitChan, tasksQueue, resultsChan)
	}

	errorCount := 0
	for i := 0; i < tasksCount; i++ {
		if result := <-resultsChan; result == nil {
			continue
		}
		errorCount++
		if errorCount == M {
			close(quitChan)
			return fmt.Errorf("errors limit was reached (%d)", errorCount)
		}
	}
	return nil
}

func tasksToQueue(quitChan <-chan bool, tasks []Task, tasksQueue chan<- Task) {
	for _, task := range tasks {
		select {
		case <-quitChan:
			close(tasksQueue)
			return
		default:
			tasksQueue <- task
		}
	}
	close(tasksQueue)
}

func worker(quitChan <-chan bool, tasksQueue <-chan Task, resultChan chan<- error) {
	for task := range tasksQueue {
		select {
		case <-quitChan:
			return
		default:
			resultChan <- task()
		}
	}
}
