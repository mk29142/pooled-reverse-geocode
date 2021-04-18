# pooled-reverse-geocode

This project takes in a set of latitudes and longitudes and returns the post code.
It used [Mapbox](https://docs.mapbox.com/api/search/geocoding/) for looking up the postcode.
You will need a api-token to run the program.

input:
```json
{ "lat": <float64>, "lng": <float64> }
```

output:
```json
{ "lat": <float64>, "lng": <float64>, "postcode": <string> }
```

## Setup
```bash
make init
```

```bash
make build
```

How to run:

```bash
 cat coordinates.txt | ./pooled-reverse-geocode "api token" "pool size flag" > output.txt
```

## Running tests

unit tests:
```bash
make test-unit
```

benchmark tests:
```bash
make test-benchmark
```

## Design decisions
 * I have decoupled the logic of actually calling the API and getting the results from the worker. 
 This will make it more extensible in the future if we needed to add different tasks that process 
 the input differently. This should just mean abstracting it behind an interface.
 
 * The worker is only responsible for possessing the task by calling `task.Process()`
 
 * I haven't added a system test at the moment. For a system test, it would mean using something like 
   [direnv](https://direnv.net/) to set a profile file with a valid api-token so that the system tests will
   test the full end-to-end flow that will actually call the api. 
 
 * The benchmark test is a bit flakey due to scheduling. Ideally I would take the average 
 of multiple runs rather than just a single run to reduce flakey results when it comes to 
 scheduling go routines.
 
 * In a production setting I would create a logger and pass in the context everywhere aswell. This would
   allow logging with structured arguments that would store the input lat and longs. 

## Notes
* I don't have much experience with concurrency in go in a production environment 
so this may not be the most idiomatic go.

* This project uses [ginkgo](https://github.com/onsi/ginkgo) for tests.
* This project used [counterfeiter](https://github.com/maxbrunsfeld/counterfeiter) to generate mocks for 
  any interfaces annotated with the following:
  ```golang
  //counterfeiter:generate -o internal/fake_client.go . Client
  type Client interface {
    Postcode(coordinates domain.Coordinates) (client.LatLongPostcode, error)
  }
  ```




