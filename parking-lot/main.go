package main

func main(){
	lot := GetParkingLot()

	level1Spots := []*ParkingSpot{
		NewParkingSpot(1, CarType),
		NewParkingSpot(2, MotorcycleType),
		NewParkingSpot(3, TruckType),
	}

	level1 := NewLevel(1, level1Spots)
	lot.AddLevel(level1)

	car := NewCar("CAR-123")
	bike := NewMotorcycle("BIKE-111")
	truck := NewTruck("TRUCK-999")

	lot.ParkVehicle(car)
	lot.ParkVehicle(bike)
	lot.ParkVehicle(truck)

	lot.ShowAvailability()

	
	lot.UnparkVehicle("CAR-123")

	lot.ShowAvailability()

}