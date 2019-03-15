package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Service struct {
	Delivery   string `json:"delivery"`
	TrackingID string `json:"trackingID"`
}

type Configuration struct {
	Services []Service `json:"services"`
}

func Load(filename string) Configuration {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Print("Error")
	}
	//defer file.Close()
	decoder := json.NewDecoder(file)
	configuration := Configuration{}
	err = decoder.Decode(&configuration)
	if err != nil {
		fmt.Println("error:", err)
	}
	return configuration
}
