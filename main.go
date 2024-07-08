package main

import (
	"fmt"
	"time"
)

type waterlevel struct {
	ID     string  `json:"id"`
	StationName  string  `json:"station_name"`
	StationId string  `json:"station_id"`
	recordDatetime  time.Time `json:"record_datetime"`
}

func main() {
	fmt.Println("WaterLevel")
}