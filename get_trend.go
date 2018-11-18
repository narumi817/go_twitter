package main

import (
	"fmt"
	"log"
	"net/url"

	. "./keys"
)

const (
	request_url = "https://api.twitter.com/1.1/trends/place.json"
	woeid       = 1118370
	// request_url = "https://api.twitter.com/1.1/trends/available.json"
)

func main() {

	v := url.Values{}

	api := GetTwitterApi()
	// trendsLocations, err := api.GetTrendsAvailableLocations(v)
	// if err != nil {
	// 	// エラーだったらログ出力
	// 	log.Fatal(err)
	// }

	// for _, trendLocation := range trendsLocations {
	// 	fmt.Println("country : ", trendLocation)
	// }

	trends, err := api.GetTrendsByPlace(woeid, v)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("trend : ", trends)
	// fmt.Printf("trend : %T", trends)
}
