package domain

type TaskError struct {
  Lat float64
  Lng float64
  Err error
}

func NewTaskError(lat, long float64, err error) TaskError {
  return TaskError{
    Lat: lat,
    Lng: long,
    Err: err,
  }
}

func (e TaskError) Error() string {
  return e.Err.Error()
}

func (e TaskError) Latitude() float64 {
  return e.Latitude()
}

func (e TaskError) Longitude() float64 {
  return e.Longitude()
}

