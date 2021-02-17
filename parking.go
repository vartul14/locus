package main

import (
	"time"
	"fmt"
)

type ParkingLotI interface {
	Park(vehicleInfo VehicleInfo, t time.Time, parkingLotID int) bool
	Exit(vehicleInfo VehicleInfo, t time.Time, parkingLotID int) (int)

}

type parkingLotImpl struct {
	ParkingLot map[int]ParkingLot  //parkinglot id,
	RateCard map[VehicleType][]TimeWindow
}

func (i *parkingLotImpl) Init(parkingLotID int, name string)  {

	pLot := make(map[int]ParkingLot)
	pLotData := ParkingLot {
		Name: name,
		Slots: map[VehicleType]int {
			Car: 1,
			Bike: 1,
		},
	}

	pLot[parkingLotID] = pLotData
	(*i).ParkingLot = pLot

	w1 := TimeWindow{
		StartTime: 0,
		EndTime: 2,
		Cost: 5.0,
	}
	w2 := TimeWindow{
		StartTime: 2,
		EndTime: 5,
		Cost: 10.0,
	}
	l := []TimeWindow {
		w1, w2,
	}
	rCard := map[VehicleType][]TimeWindow {
		Car: l,
	}

	(*i).RateCard = rCard
}

func (i *parkingLotImpl) Park(vehicleInfo VehicleInfo, t time.Time, parkingLotID int) bool {
	//check if slots available
	isAvailable := i.checkIfSlotsAvailable(vehicleInfo.VehicleType, parkingLotID)
	if !isAvailable {
		fmt.Println("No parking slots available")
		return false
	}

	//park the vehicle with the entry time stamp
	i.parkVehicle(vehicleInfo, t, parkingLotID)

	//Reduce available slots
	i.reduceAvailabeSlots(vehicleInfo.VehicleType, parkingLotID)

	return true
}

func (i *parkingLotImpl) Exit(vehicleInfo VehicleInfo, t time.Time, parkingLotID int) float64 {
	//remove the vehicle from the parking lot
	pData, _ := i.exitVehicle(vehicleInfo, parkingLotID)

	//inc available slots
	i.increaseAvailableSlots(vehicleInfo.VehicleType, parkingLotID)

	fare := i.computeFare(pData, t)

	return fare
}



func (i *parkingLotImpl)checkIfSlotsAvailable(vehicleType VehicleType, parkingLotID int) bool {
	parkingLot := (*i).ParkingLot[parkingLotID]
	if parkingLot.Slots == nil {
		parkingLot.Slots = make(map[VehicleType]int)
	}

	if parkingLot.Slots[vehicleType] > 0 {
		return true
	}

	return false
}

func (i *parkingLotImpl) parkVehicle(vehicleInfo VehicleInfo, t time.Time, parkingLotID int) {
	parkingLot := (*i).ParkingLot[parkingLotID]

	if parkingLot.ParkingInfo == nil {
		pInfo := make(map[string]ParkingData)
		parkingLot.ParkingInfo = pInfo 
	}

	parkingData := ParkingData {
		EntryTime: t,
		VehicleInfo: vehicleInfo,
	}
	parkingLot.ParkingInfo[vehicleInfo.RegdNo] = parkingData
	(*i).ParkingLot[parkingLotID] = parkingLot
}

func (i *parkingLotImpl) reduceAvailabeSlots(vehicleType VehicleType, parkingLotID int) {
	parkingLot := (*i).ParkingLot[parkingLotID]

	slots := parkingLot.Slots[vehicleType]
	parkingLot.Slots[vehicleType] = slots - 1
}

func (i *parkingLotImpl) exitVehicle(vehicleInfo VehicleInfo, parkingLotID int) (ParkingData, error) {
	parkingLot := (*i).ParkingLot[parkingLotID]

	pInfo := parkingLot.ParkingInfo

	pData, ok := pInfo[vehicleInfo.RegdNo]
	if !ok {
		fmt.Println("Vehicle not present in parking lot")
		return ParkingData{}, fmt.Errorf("Vehicle not present in parking lot")
	}

	delete(pInfo, vehicleInfo.RegdNo)
	return pData, nil
}

func (i *parkingLotImpl) computeFare(parkingData ParkingData, exTime time.Time) float64 {
	rateCard := (*i).RateCard
	vType := parkingData.VehicleInfo.VehicleType

	eTime := parkingData.EntryTime
	rateCardForVehicle := rateCard[vType]
	tDiff := exTime.Sub(eTime).Hours()

	for _, val := range rateCardForVehicle {
		if tDiff <= val.EndTime {
			
			return val.Cost
		}
	}
	return 0.0
	
	
}

func (i *parkingLotImpl) increaseAvailableSlots(vehicleType VehicleType, parkingLotID int) {
	parkingLot := (*i).ParkingLot[parkingLotID]

	slots := parkingLot.Slots[vehicleType]
	parkingLot.Slots[vehicleType] = slots + 1
}

