package main

import (
	"flag"
	"fmt"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
	"os"
	"strconv"
	"time"
)

// RuneScapeUsernameInput is the user to export to Prometheus
var RuneScapeUsernameInput = flag.String("runescape-username", "", "")
var WebPortInput = flag.Int("port", 8080, "")
var UpdatePeriod = flag.Duration("update-period", time.Minute*5, "")
var RuneScapeUsername = ""

func main() {
	flag.Parse()
	if RuneScapeUsernameInput == nil || *RuneScapeUsernameInput == "" {
		fmt.Println("You must specify RunescapeUsername with --runescape-username=")
		os.Exit(2)
	}
	RuneScapeUsername = *RuneScapeUsernameInput
	if WebPortInput == nil || *WebPortInput == 0 {
		fmt.Println("You must specify WebPort with --port")
		os.Exit(2)
	}
	var WebPort = ":" + strconv.Itoa(*WebPortInput)
	if UpdatePeriod == nil || *UpdatePeriod == 0 {
		fmt.Println("You must specify UpdatePeriod with --updateperiod")
	}

	updateMetrics()
	http.Handle("/metrics", promhttp.Handler())
	err := http.ListenAndServe(WebPort, nil)
	if err != nil {
		return
	}
}
