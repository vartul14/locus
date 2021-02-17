package main

import(
	"testing"
	"time"
	"fmt"
)

func TestPark (t *testing.T) {
	impl := parkingLotImpl{}

	vehicleInfo1 := VehicleInfo{
		RegdNo: "KA08J7899",
		VehicleType: Car,
	}

	vehicleInfo2 := VehicleInfo{
		RegdNo: "KA08J7800",
		VehicleType: Car,
	}

	impl.Init(1, "p1")
	status := impl.Park(vehicleInfo1, time.Now(), 1)
	if status != true {
		t.Errorf("Erro in parking")
	}
	status2 := impl.Park(vehicleInfo2, time.Now(), 1)
	if status2 != false {
		t.Errorf("Error should be thrown")
	}

	fmt.Println(impl.ParkingLot)

	cost := impl.Exit(vehicleInfo1, time.Now(), 1)
	if cost != 5.0{
		t.Errorf("Error  in fare")
	}
}