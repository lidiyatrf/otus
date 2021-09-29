package hw05parallelexecution

import (
	"errors"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

type Task func() error

func Run(tasks []Task, n, m int) (resultErr error) {
	tasksCh := make(chan Task, n)
	lastTaskID := n
	for i := 0; i < lastTaskID; i++ {
		tasksCh <- tasks[i]
	}

	errorsCh := make(chan error)
	terminatedCh := make(chan struct{})

	for i := 0; i < n; i++ {
		go doTask(tasksCh, terminatedCh, errorsCh)
	}

	for {
		select {
		case err := <-errorsCh:
			if m < 0 || lastTaskID > len(tasks) {
				break
			}
			if err != nil {
				m--
				if m == 0 {
					resultErr = ErrErrorsLimitExceeded
					close(tasksCh)
				}
			}
			if lastTaskID != len(tasks) && resultErr == nil {
				tasksCh <- tasks[lastTaskID]
				lastTaskID++
				break
			}
			if lastTaskID == len(tasks) && resultErr == nil {
				lastTaskID++
				close(tasksCh)
			}
		case <-terminatedCh:
			n--
			if n == 0 {
				return resultErr
			}
		}
	}
}

func doTask(tasksCh chan Task, terminatedCh chan struct{}, errorsCh chan error) {
	for {
		task, ok := <-tasksCh
		if !ok {
			terminatedCh <- struct{}{}
			return
		}
		errorsCh <- task()
	}
}
