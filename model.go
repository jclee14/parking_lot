package main

type parkingSlot struct {
	parkedCar *car
}

type car struct {
	registrationNumber string
	color              string
}
