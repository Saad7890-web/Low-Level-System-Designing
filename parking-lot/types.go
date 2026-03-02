package main

type VehicleType int

const (
	MotorcycleType VehicleType = iota
	CarType
	TruckType
)

func (v VehicleType) String() string {
	switch v {
	case MotorcycleType:
		return "Motorcycle"

	case CarType:
		return  "Car"

	case TruckType:
		return "Truck"

	default:
		return "Unknown"
	}
}