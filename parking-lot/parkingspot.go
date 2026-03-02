package main

import "sync"

// The ParkingSpot class represents an individual parking spot and tracks the availability and the parked vehicle.

type ParkingSpot struct {
	ID int
	AllowedType VehicleType
	Vehicle Vehicle
	mutex sync.Mutex
}

func NewParkingSpot(id int, vType VehicleType) *ParkingSpot {
	return &ParkingSpot{
		ID: id,
		AllowedType: vType,
	}
}


func (p *ParkingSpot) IsAvailable() bool {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	return p.Vehicle == nil
}

func (p *ParkingSpot) Park(vehicle Vehicle) bool {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	if p.Vehicle != nil || vehicle.GetType() != p.AllowedType {
		return false
	}

	p.Vehicle = vehicle
	return true
}


func (p *ParkingSpot) Unpark(){
	p.mutex.Lock()
	defer p.mutex.Unlock()

	p.Vehicle = nil
}