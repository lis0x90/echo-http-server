package main

import (
	"flag"
	"fmt"
	"net/http"
	"time"
)

func constructHandler(sleepTime time.Duration) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if sleepTime > 0 {
			time.Sleep(sleepTime)
		}
		fmt.Fprintf(w, "Path: %s", r.URL.Path[1:])
	}
}

func main() {
	var bindAddress string
	flag.StringVar(&bindAddress, "bind", "localhost:8080", "Bind adress and port")

	var sleepTime time.Duration
	flag.DurationVar(&sleepTime, "sleep-time", time.Duration(0),
		"Sleep specified duration before send response, therefore we can emulate slow servers")
	flag.Parse()

	fmt.Println("Start http server on:", bindAddress, ", sleep time:", sleepTime)
	http.HandleFunc("/", constructHandler(sleepTime))
	http.ListenAndServe(bindAddress, nil)
}
