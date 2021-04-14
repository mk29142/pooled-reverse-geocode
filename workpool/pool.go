package workpool

import (
  "github.com/mk29142/pooled-reverse-geocode/task"
  "sync"
)

type Pool struct {
  Tasks   []task.Task

  concurrency   int
  input         chan task.Task
  output        chan task.CoordinatesWithPostcode
  errors        chan error
  wg            sync.WaitGroup
}

func New(tasks []task.Task, concurrency int) Pool {
  return Pool{
    Tasks:       tasks,
    concurrency: concurrency,
    input:       make(chan task.Task),
    output:      make(chan task.CoordinatesWithPostcode),
    errors:      make(chan error),
  }
}

func (p Pool) Output() <-chan task.CoordinatesWithPostcode {
  return p.output
}

func (p Pool) Errors() <-chan error {
  return p.errors
}

func (p Pool) Run() {
  for i := 1; i <= p.concurrency; i++ {
    worker := NewWork(p.input, p.output, p.errors)
    worker.Start(&p.wg)
  }

  for i := range p.Tasks {
    p.input <- p.Tasks[i]
  }
  close(p.input)
}

func (p Pool) Cleanup() {
  go func() {
    p.wg.Wait()
    close(p.output)
    close(p.errors)
  }()
}

