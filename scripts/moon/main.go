package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

var (
	//moonPhasesAll = []string{
	//" ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ",
	//}

	moonPhases = map[string]string{
		"New":             " ",
		"Waxing Crescent": " ",
		"First Quarter":   " ",
		"Waxing Gibbous":  " ",
		"Full":            " ",
		"Waning Gibbous":  " ",
		"Last Quarter":    " ",
		"Waning Crescent": " ",
	}
)

const (
	oldDataWaring = "!"
	fileName      = "last-update.txt"
	url           = "https://moon-phase.p.rapidapi.com/plain-text"
	host          = "moon-phase.p.rapidapi.com"
)

func main() {
	currentTime := time.Now()
	hour := currentTime.Hour()
	if !(hour >= 18 || hour <= 8) {
		return
	}

	req, err := http.NewRequest("GET", url, nil)
	req.Header.Add("X-RapidAPI-Key", os.Getenv("RAPIDAPI_KEY"))
	req.Header.Add("X-RapidAPI-Host", host)

	var waring string
	var phase string

	res, err := http.DefaultClient.Do(req)

	if err != nil || res.StatusCode != http.StatusOK {
		waring = oldDataWaring
		//if request error: read last update from file
		content, err := os.ReadFile(fileName)
		if err != nil {
			fmt.Println("M: err: file read")
		}
		phase = string(content)

	} else {
		defer res.Body.Close()
		//if res.StatusCode != http.StatusOK {
		//fmt.Println(" 󰒏 ", res.StatusCode)
		//}

		body, err := io.ReadAll(res.Body)
		if err != nil {
			fmt.Println("M: err: body read")
		}

		phase = string(body)
		out, _ := os.Create(fileName)
		defer out.Close()
		out.Write(body)
	}

	fmt.Printf("%s%s", waring, moonPhases[phase])
}
