package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
)

type Icon map[int]struct {
	iconDay, iconNight, color string
}

type Weather struct {
	CurrentWeather struct {
		Temperature float64 `json:"temperature"`
		Weathercode int     `json:"weathercode"`
		IsDay       int     `json:"is_day"`
	} `json:"current_weather"`
}

// Lviv
// 49.8383  24.0232
const (
	// some hardcode here :)
	weatherApi = "https://api.open-meteo.com/v1/forecast?latitude=%f&longitude=%f&current_weather=true&timezone=Europe%%2FKiev"
	fileName   = "last-update.json"

	s2d_bg0     = "^c#1d2021^"
	s2d_fg0     = "^c#fbf1c7^"
	s2d_fg1     = "^c#ebdbb2^"
	s2d_red0    = "^c#cc241d^"
	s2d_red1    = "^c#fb3943^"
	s2d_green0  = "^c#98971a^"
	s2d_green1  = "^c#b8bb26^"
	s2d_yellow0 = "^c#d79921^"
	s2d_yellow1 = "^c#fabd2f^"
	s2d_blue0   = "^c#558588^"
	s2d_blue1   = "^c#83a598^"
	s2d_purple0 = "^c#b16286^"
	s2d_purple1 = "^c#d3869b^"
	s2d_aqua0   = "^c#689d6a^"
	s2d_aqua1   = "^c#8ec07c^"
	s2d_orange0 = "^c#d65d0e^"
	s2d_orange1 = "^c#fe8019^"
	s2d_reset   = "^d^"
)

var (
	icons = Icon{
		//wc day   night color
		0:  {" ", " ", s2d_yellow1},
		1:  {" ", " ", s2d_yellow1},
		2:  {" ", " ", s2d_fg1},
		3:  {"󰖐 ", "󰖐 ", s2d_fg1},
		45: {" ", " ", s2d_fg1},
		48: {" ", " ", s2d_fg1},
		51: {" ", " ", s2d_blue1},
		53: {" ", " ", s2d_blue1},
		55: {" ", " ", s2d_blue1},
		56: {" ", " ", s2d_blue1},
		57: {" ", " ", s2d_blue1},
		61: {" ", " ", s2d_blue1},
		63: {" ", " ", s2d_blue1},
		65: {" ", " ", s2d_blue1},
		66: {"󰙿 ", "󰙿 ", s2d_fg0},
		67: {"󰙿 ", "󰙿 ", s2d_fg0},
		71: {" ", " ", s2d_fg0},
		73: {" ", " ", s2d_fg0},
		75: {" ", " ", s2d_fg0},
		77: {" ", " ", s2d_fg0},
		80: {" ", " ", s2d_blue1},
		81: {" ", " ", s2d_blue1},
		82: {"", "", s2d_blue1},
		85: {"󰜗 ", "󰜗 ", s2d_fg0},
		86: {" ", " ", s2d_fg0},
		95: {" ", " ", s2d_purple1},
		96: {" ", " ", s2d_purple1},
		99: {" ", " ", s2d_purple1},
	}

	hot   = 28.0
	cold  = 15.0
	minus = 0.0

	oldDataWarning = "!"
)

func main() {
	//check latitude and longitude in arguments
	if len(os.Args) < 3 {
		fmt.Println("W: err: args")
		return
	}
	latitude, err := strconv.ParseFloat(os.Args[1], 32)
	if err != nil {
		fmt.Println("W: err: type")
		return
	}
	longitude, err := strconv.ParseFloat(os.Args[2], 32)
	if err != nil {
		fmt.Println("W: err: type")
		return
	}

	var weather Weather
	var warning string

	// not the cleverest way to format URL
	weatherURL := fmt.Sprintf(weatherApi, latitude, longitude)
	// make request
	resp, err := http.Get(weatherURL)
	if err != nil {
		warning = oldDataWarning
		//if request error: read last update from file
		content, err := os.ReadFile(fileName)
		if err != nil {
			fmt.Println("W: err: file read")
		}
		err = json.Unmarshal(content, &weather)
		if err != nil {
			fmt.Println("W: err: unmarshal")
		}
	} else {
		// Decode response
		err = json.NewDecoder(resp.Body).Decode(&weather)
		if err != nil {
			fmt.Println("W: err: decode")
		}
		// Save JSON content to file
		out, _ := os.Create(fileName)
		defer out.Close()
		json.NewEncoder(out).Encode(&weather)
		defer resp.Body.Close()
	}

	var textColor string
	switch {
	case weather.CurrentWeather.Temperature >= hot:
		textColor = s2d_orange1
	case weather.CurrentWeather.Temperature >= cold:
		textColor = s2d_fg0
	case weather.CurrentWeather.Temperature >= minus:
		textColor = s2d_blue1
	default:
		textColor = s2d_blue0
	}

	var icon string
	var iconColor string
	if weather.CurrentWeather.IsDay == 1 {
		icon = icons[weather.CurrentWeather.Weathercode].iconDay
		iconColor = icons[weather.CurrentWeather.Weathercode].color
	} else {
		icon = icons[weather.CurrentWeather.Weathercode].iconNight
		iconColor = s2d_blue0
	}
	fmt.Printf("%s%s%s%s%.1f°C%s", warning, iconColor, icon, textColor, weather.CurrentWeather.Temperature, s2d_reset)
}
