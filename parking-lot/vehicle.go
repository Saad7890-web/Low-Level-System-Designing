package main

// The Vehicle class is an abstract base class for different types of vehicles. It is extended by Car, Motorcycle, and Truck classes.
type Vehicle interface {
	GetLicensePlate() string
	GetType() VehicleType
}

type BaseVehicle struct {
	LicensePlate string
	Type VehicleType
}

func (v *BaseVehicle) GetLicensePlate() string  {
	return v.LicensePlate
}

func (v *BaseVehicle) GetType() VehicleType {
	return v.Type
}

func NewCar(plate string) Vehicle{
	return &BaseVehicle{
		LicensePlate: plate,
		Type: CarType,
	}
}

func NewMotorcycle(plate string) Vehicle{
	return &BaseVehicle{
		LicensePlate: plate,
		Type: MotorcycleType,
	}
}

func NewTruck(plate string) Vehicle{
	return &BaseVehicle{
		LicensePlate: plate,
		Type: TruckType,
	}
}