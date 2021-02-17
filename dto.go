package main

import (
	"time"
)

type ParkingLot struct {
	Name string
	Slots map[VehicleType]int    //map[VehicleType]map[FloorID]int
	ParkingInfo map[string]ParkingData
}

type VehicleInfo struct {
	RegdNo string
	VehicleType VehicleType
}

type ParkingData struct {
	EntryTime time.Time
	VehicleInfo VehicleInfo
}

type TimeWindow struct {
	StartTime float64
	EndTime float64
	Cost float64
}
