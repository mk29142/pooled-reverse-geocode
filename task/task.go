package task

import (
  "fmt"
  "github.com/mk29142/pooled-reverse-geocode/client"
  "github.com/mk29142/pooled-reverse-geocode/domain"
)

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 -generate

//counterfeiter:generate -o internal/fake_client.go . Client
type Client interface {
  Postcode(coordinates domain.LatLong) (client.LatLongPostcode, error)
}

type Task struct {
  LatLong  domain.LatLong
  GeoCoder Client
}

type Result struct {
  PostCode client.LatLongPostcode
}

func NewTask(latLong domain.LatLong, geocoder Client) Task {
  return Task{
    LatLong:  latLong,
    GeoCoder: geocoder,
  }
}

func (task Task) Process() (Result, error) {
  res, err := task.GeoCoder.Postcode(task.LatLong)
  if err != nil {
    return Result{},
    domain.NewTaskError(task.LatLong.Latitude,
      task.LatLong.Longitude,
      fmt.Errorf("failure process task: %w", err))
    }

  return Result{PostCode:res}, nil
}