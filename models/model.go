package models

type ParkingSlot struct {
	ParkedCar *Car
}

type Car struct {
	RegistrationNumber string
	Color              string
}
