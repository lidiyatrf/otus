package hw05parallelexecution

import (
	"errors"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

type Task func() error

func Run(tasks []Task, n, m int) (resultErr error) {
	ignoreErrors := m <= 0

	tasksCh := make(chan Task, n)
	errorsCh := make(chan error)
	terminatedCh := make(chan struct{})

	for i := 0; i < n; i++ {
		tasksCh <- tasks[i]
		go doTask(tasksCh, terminatedCh, errorsCh)
	}

	lastTaskID := n

	for {
		select {
		case err := <-errorsCh:
			if !ignoreErrors && m < 0 || lastTaskID > len(tasks) || resultErr != nil {
				break
			}
			if err != nil {
				m--
				if m == 0 {
					resultErr = ErrErrorsLimitExceeded
					close(tasksCh)
					break
				}
			}
			if lastTaskID != len(tasks) {
				tasksCh <- tasks[lastTaskID]
				lastTaskID++
				break
			}
			lastTaskID++
			close(tasksCh)
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
