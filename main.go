package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/mk29142/pooled-reverse-geocode/client"
	"github.com/mk29142/pooled-reverse-geocode/domain"
	"github.com/mk29142/pooled-reverse-geocode/task"
	"github.com/mk29142/pooled-reverse-geocode/workpool"
	"net/http"
	"os"
	"strconv"
)

func main() {

	//TODO: make only api-token mandatory
	argLength := len(os.Args[1:])
	if argLength < 2 {
		fmt.Fprintln(os.Stderr, "must provide api token and pool size")
		os.Exit(2)
	}

	poolSizeArg := os.Args[argLength]
	poolSize, err := strconv.Atoi(poolSizeArg)
	if err != nil {
		fmt.Println("must provide valid pool size")
		os.Exit(2)
	}

	apiToken := os.Args[argLength-1]

	client := client.New(apiToken, &http.Client{})

	var tasks []task.Task
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		var latlong domain.LatLong
		json.Unmarshal([]byte(scanner.Text()), &latlong)
		tasks = append(tasks, task.NewTask(latlong, client))
	}

	pool := workpool.New(tasks, poolSize)

	go func() {
		// TODO: return as json string
		for res := range pool.Output() {
			fmt.Fprintln(os.Stdout, domain.Postcode{
				Latitude:  res.PostCode.Latitude,
				Longitude: res.PostCode.Longitude,
				Postcode:  res.PostCode.Postcode,
			})
		}
	}()

	go func() {
		for err := range pool.Errors() {
			fmt.Fprintln(os.Stderr,  err)
		}
	}()

	pool.Run()
  pool.Cleanup()
}
