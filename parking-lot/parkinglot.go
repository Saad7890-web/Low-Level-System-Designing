package main

import (
	"fmt"
	"sync"
)

type ParkingLot struct {
	Levels []*Level
}

var instance *ParkingLot
var one sync.Once


func GetParkingLot() *ParkingLot {
	one.Do(func ()  {
		instance = &ParkingLot{}
	})

	return instance
}

func (p *ParkingLot) AddLevel(level *Level){
	p.Levels = append(p.Levels, level)
}

func (p *ParkingLot) ParkVehicle(vehicle Vehicle) {
	for _, level := range p.Levels {
		spot := level.ParkVehicle(vehicle)
		if spot != nil {
			fmt.Printf("Vehicle %s parked at Level %d Spot %d\n",
				vehicle.GetLicensePlate(),
				level.FloorNumber,
				spot.ID)
			return
		}
	}
	fmt.Println("No available spot for vehicle:", vehicle.GetLicensePlate())
}

func (p *ParkingLot) UnparkVehicle(plate string) {
	for _, level := range p.Levels {
		if level.UnparkVehicle(plate) {
			fmt.Println("Vehicle", plate, "unparked")
			return
		}
	}
	fmt.Println("Vehicle not found:", plate)
}

func (p *ParkingLot) ShowAvailability() {
	for _, level := range p.Levels {
		fmt.Printf("Level %d Available Spots: %d\n",
			level.FloorNumber,
			level.AvailableSpots())
	}
}