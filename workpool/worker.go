package workpool

import (
  "sync"
)

type Worker struct {
  Tasks   chan Task
  Outputs chan CoordinatesWithPostcode
  Errors  chan error
}

func NewWorker(tasks chan Task, output chan CoordinatesWithPostcode, errors chan error) Worker {
  return Worker{
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


