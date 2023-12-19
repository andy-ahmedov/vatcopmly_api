package main

import (
	"log"
	"time"

	"github.com/andy-ahmedov/vatcomply_api/vatcomply"
)

func main() {
	vatcomplyClient, err := vatcomply.NewClient(time.Second * 10)
	if err != nil {
		log.Fatal(err)
	}

	rates, err := vatcomplyClient.GetBaseRates("CZK")
	if err != nil {
		log.Fatal(err)
	}
	rates.Info()

	geo, err := vatcomplyClient.GetGeolocation()
	if err != nil {
		log.Fatal(err)
	}
	geo.Info()

	date, err := vatcomplyClient.GetDateeRates("2000-04-05")
	if err != nil {
		log.Fatal(err)
	}
	date.Info()

	late, err := vatcomplyClient.GetLatestRates()
	if err != nil {
		log.Fatal(err)
	}
	late.Info()
}
