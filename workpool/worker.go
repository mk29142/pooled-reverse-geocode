package workpool

import (
  "github.com/mk29142/pooled-reverse-geocode/task"
  "sync"
)

type Worker struct {
  Tasks   chan task.Task
  Outputs chan task.CoordinatesWithPostcode
  Errors  chan error
}

func NewWork(tasks chan task.Task, output chan task.CoordinatesWithPostcode, errors chan error) *Worker {
  return &Worker{
    Tasks:     tasks,
    Outputs:   output,
    Errors:    errors,
  }
}

func (wr Worker) Start(wg *sync.WaitGroup) {
  wg.Add(1)

  go func() {
    defer wg.Done()
    for task := range wr.Tasks {
      res, err := task.Process()
      if err != nil {
        wr.Errors <- err
        continue
      }

      wr.Outputs <- res
    }
  }()
}


