package client

import (
  "encoding/json"
  "fmt"
  "io"
  "net/http"
  "strings"

  "github.com/mk29142/pooled-reverse-geocode/domain"
)

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 -generate

type LatLongPostcode struct {
  Latitude float64
  Longitude float64
  Postcode string
}

type location struct {
  Features []feature `json:"Features"`
}

type feature struct {
    Text string `json:"text"`
}

//counterfeiter:generate -o internal/fake_client.go . client
type client interface {
  Do(*http.Request) (*http.Response, error)
}

type Client struct {
  ApiToken    string
  Client      client
}

func New(apiToken string, client client) Client {
  return Client{
    ApiToken: apiToken,
    Client: client,
  }
}

func (c Client) Postcode(coordinates domain.LatLong) (LatLongPostcode, error) {
   url := fmt.Sprintf("https://api.mapbox.com/geocoding/v5/mapbox.places/%f,%f.json?types=postcode&limit=1&access_token=%s", coordinates.Longitude, coordinates.Latitude, c.ApiToken)

  request, err := http.NewRequest(http.MethodGet, url, nil)
  if err != nil {
    return LatLongPostcode{}, fmt.Errorf("create request: %w", err)
  }

  response, err := c.Client.Do(request)
  if err != nil {
    return LatLongPostcode{}, fmt.Errorf("get request: %w", err)
  }

  if response.StatusCode != http.StatusOK {
    return LatLongPostcode{}, fmt.Errorf(`unexpected status code "%s"`, strings.ToLower(response.Status))
  }

  var loc location
  if err := json.NewDecoder(response.Body).Decode(&loc); err != nil {
    return LatLongPostcode{}, fmt.Errorf("decode response body: %w", err)
  }

  postcode := loc.Features[0].Text

  closeIgnoreErr(response.Body)
  return LatLongPostcode{
    Latitude:  coordinates.Latitude,
    Longitude: coordinates.Longitude,
    Postcode:  postcode,
  }, nil
}

func closeIgnoreErr(c io.Closer) {
  _ = c.Close()
}
