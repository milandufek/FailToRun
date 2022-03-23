package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os/exec"
	"strings"
	"time"
)

/*
Simple program.
Get URL, if response is lower/higher that specified, then run command on your system.
*/

func main() {

	// parse args
	var url = flag.String("u", "www.google.com", "URL to check.")
	var command = flag.String("c", "", "Command to run.")
	var onBackGround = flag.Int("b", 1, "Run command at background.")
	var maxRepeat = flag.Int("m", 0, "Maximum number of repeats (0 = unlimited).")
	var period = flag.Int("p", 1, "Repeat every N seconds.")
	var requestTimeout = flag.Int("t", 2, "Request timeout in seconds.")
	var maxExpectedResponse = flag.Int("r", http.StatusOK, "Maximum response code.")
	var negationFlag = flag.Int("n", 0, "Negator (run at success condition).")
	flag.Parse()

	var reqCount = 1

	if *url == "" {
		log.Fatal("URL is empty")
	}

	// proceed request(s)
	for true {
		var runAtSuccess = true
		if *negationFlag == 1 {
			runAtSuccess = !runAtSuccess
		}

		status := httpGetStatus(*url, *requestTimeout, *maxExpectedResponse)

		// Run on false status unless negation flag is set to true. (XOR)
		if status != runAtSuccess {
			runCmd(*command, *onBackGround)
			break
		}

		if *period < 1 || *maxRepeat == reqCount {
			break
		}

		time.Sleep(time.Duration(*period) * time.Second)

		if *maxRepeat != 0 {
			reqCount++
		}
	}
}

func httpGetStatus(url string, timeout int, maxExpectedResponse int) bool {

	client := http.Client{Timeout: time.Duration(timeout) * time.Second}

	if !strings.HasPrefix(url, "http://") || !strings.HasPrefix(url, "https://") {
		url = "https://" + url
	}

	log.Printf("Requesting URL: %s", url)
	resp, err := client.Get(url)

	if err != nil {
		log.Printf("HTTP error [%s]\n", err)
		return false
	}

	defer resp.Body.Close()

	if resp.StatusCode > maxExpectedResponse {
		log.Printf("HTTP request failed [%s]\n", resp.Status)
		return false
	}

	log.Printf("Response [%d]", resp.StatusCode)

	return true
}

func runCmd(input string, onBackGround int) {

	var cmd = strings.TrimSpace(input)

	slices := strings.Split(cmd, " ")
	bin := slices[0]
	args := slices[1:]

	if bin == "" {
		log.Print("No binary specified to run")
		return
	}

	if onBackGround == 1 {
		log.Printf("Running command [%s] at background", cmd)
		err := exec.Command(bin, args...).Start()
		if err != nil {
			log.Println("Command failed")
		}

	} else {
		log.Printf("Running command [%s]", cmd)
		out, err := exec.Command(bin, args...).Output()
		if err != nil {
			log.Println("Command failed")
		}
		fmt.Printf("%s\n", out)
	}
}
