package main

//Each level:

//Has multiple parking spots

//Can park/unpark vehicles

//Checks availability

type Level struct {
	FloorNumber int
	Spots []*ParkingSpot
}

func NewLevel(floor int, spots []*ParkingSpot) *Level {
	return &Level{
		FloorNumber: floor,
		Spots: spots,
	}
}

func (l *Level) ParkVehicle(vehicle Vehicle) *ParkingSpot {
	for _,spot := range l.Spots {
		if spot.IsAvailable() && spot.AllowedType == vehicle.GetType(){
			if spot.Park(vehicle){
				return spot
			}

		}

	}
	return nil
}

func (l *Level) UnparkVehicle(plate string) bool {
	for _, spot := range l.Spots {
		if !spot.IsAvailable() &&
		spot.Vehicle.GetLicensePlate() == plate {
			spot.Unpark()
			return true
		}
	}
	return false
}


func (l *Level) AvailableSpots() int {
	count := 0
	for _,spot := range l.Spots {
		if spot.IsAvailable() {
			count++
		}
	}
	return count
}
