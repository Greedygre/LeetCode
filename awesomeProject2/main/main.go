package main

import (
	"log"
	"fmt"
	"os"

	// Shortening the import reference name seems to make it a bit easier
	owm "github.com/briandowns/openweathermap"
)

var apiKey = os.Getenv("058276ece6f19db2b6580b536e22d780")
func main() {
	w, err := owm.NewForecast("5", "F", "FI", apiKey) // valid options for first parameter are "5" and "16"
	if err != nil {
		log.Fatalln(err)
	}

	w.DailyByCoordinates(&owm.Coordinates{
			Longitude: -112.07,
			Latitude: 33.45,
		},
		5, // five days forecast
	)
	fmt.Println(w)
}