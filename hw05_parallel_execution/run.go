package hw05parallelexecution

import (
	"errors"
	"sync"
	"sync/atomic"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")
var wg = &sync.WaitGroup{}

type Task func() error

// Run starts tasks in n goroutines and stops its work when receiving m errors from tasks.
func Run(tasks []Task, n, m int) (e error) {
	if n <= 0 || m <= 0 {
		return ErrErrorsLimitExceeded
	}

	var errorsCount int64
	wg.Add(n)
	tasksChan := make(chan Task, len(tasks)+1)

	for i := 0; i < len(tasks); i++ {
		tasksChan <- tasks[i]
	}
	close(tasksChan)

	for i := 0; i < n; i++ {
		go process(tasksChan, &errorsCount, m)
	}

	wg.Wait()

	if int(errorsCount) >= m {
		e = ErrErrorsLimitExceeded
	}

	return e
}

func process(tasksChan chan Task, errorsCount *int64, m int) {
	defer wg.Done()

	for task := range tasksChan {
		err := task()
		if err != nil {
			atomic.AddInt64(errorsCount, 1)
		}

		if int(atomic.LoadInt64(errorsCount)) >= m {
			return
		}
	}
}
