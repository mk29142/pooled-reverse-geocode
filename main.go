package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/mk29142/pooled-reverse-geocode/client"
	"github.com/mk29142/pooled-reverse-geocode/domain"
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

	var tasks []workpool.Task
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		var latlong domain.Coordinates
		json.Unmarshal([]byte(scanner.Text()), &latlong)
		tasks = append(tasks, workpool.NewTask(latlong, client))
	}

	pool := workpool.New(tasks, poolSize)

	go func() {
		for res := range pool.Output() {
			postcode := domain.Postcode{
				Latitude:  res.Lat,
				Longitude: res.Long,
				Postcode:  res.PostCode,
			}

			out, _ := json.Marshal(postcode)
			fmt.Fprintln(os.Stdout, string(out))
		}
	}()

	go func() {
		for err := range pool.Errors() {
			fmt.Fprintln(os.Stderr,  err)
		}
	}()

	pool.Run()
}
