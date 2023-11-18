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

	rates, err := vatcomplyClient.GetLatestRates()
	if err != nil {
		log.Fatal(err)
	}
	rates.Info()
}
