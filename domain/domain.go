package domain

type Coordinates struct {
  Latitude float64 `json:"lat"`
  Longitude float64 `json:"lng"`
}

type Postcode struct {
  Latitude  float64 `json:"lat"`
  Longitude float64 `json:"lng"`
  Postcode  string  `json:"postcode"`
}
