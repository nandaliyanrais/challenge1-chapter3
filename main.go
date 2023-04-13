package main

import (
	// "encoding/json"
	"fmt"
	"math/rand"
	"time"
)

type Weather struct {
	Water       int    `json:"Water"`
	Wind        int    `json:"Wind"`
}

type Status struct {
	WaterStatus string `json:"Status Water"`
	WindStatus  string `json:"Status Wind"`
}

func main() {

	timer := time.NewTicker(15 * time.Second)
	defer timer.Stop()

	for range timer.C {

		water := rand.Intn(100) + 1
		wind := rand.Intn(100) + 1
	
		waterStatus := "aman"
		if water >= 6 && water <= 8 {
			waterStatus = "siaga"
		} else if water > 8 {
			waterStatus = "bahaya"
		}
	
		windStatus := "aman"
		if wind >= 7 && wind <= 15 {
			windStatus = "siaga"
		} else if wind > 15 {
			windStatus = "bahaya"
		}
	
		weather := Weather{
			Water:       water,
			Wind:        wind,
		}

		status := Status{
			WaterStatus: waterStatus,
			WindStatus:  windStatus,
		}
	
		// weatherJSON, err := json.Marshal(weather)
		// if err != nil {
		// 	fmt.Println("Error marshaling weather object to JSON:", err)
		// 	continue
		// }

		// fmt.Println(string(weatherJSON))

		fmt.Printf("\n{\n\tWater : %d,\n\tWind : %d\n}\n", weather.Water, weather.Wind)

		fmt.Println("Status Water :", string(status.WaterStatus))
		fmt.Println("Status Wind  :", string(status.WindStatus))

	}
}