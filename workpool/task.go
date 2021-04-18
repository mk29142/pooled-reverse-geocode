package workpool

import (
  "fmt"
  "github.com/mk29142/pooled-reverse-geocode/client"
  "github.com/mk29142/pooled-reverse-geocode/domain"
)

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 -generate

//counterfeiter:generate -o internal/fake_client.go . Client
type Client interface {
  Postcode(coordinates domain.Coordinates) (client.LatLongPostcode, error)
}

type Task struct {
  Coordinates domain.Coordinates
  GeoCoder    Client
}

type CoordinatesWithPostcode struct {
  Lat float64
  Long float64
  PostCode string
}

func NewTask(latLong domain.Coordinates, geocoder Client) Task {
  return Task{
    Coordinates: latLong,
    GeoCoder:    geocoder,
  }
}

func (task Task) Process() (CoordinatesWithPostcode, error) {
  res, err := task.GeoCoder.Postcode(task.Coordinates)
  if err != nil {
    return CoordinatesWithPostcode{},
    domain.NewTaskError(task.Coordinates.Latitude,
      task.Coordinates.Longitude,
      fmt.Errorf("failure process task: %w", err))
    }

  return CoordinatesWithPostcode{
    Lat:      res.Latitude,
    Long:     res.Longitude,
    PostCode: res.Postcode,
  }, nil
}