package models

type ParkingSlot struct {
	ParkedCar *Car
}

type Car struct {
	RegistrationNumber string
	Color              string
}

type CommandSet string

const (
	CommandSetCreateParkingLot                  CommandSet = "create_parking_lot"
	CommandSetPark                              CommandSet = "park"
	CommandSetLeave                             CommandSet = "leave"
	CommandSetStatus                            CommandSet = "status"
	CommandSetRegistrationNumberForCarWithColor CommandSet = "registration_numbers_for_cars_with_colour"
	CommandSetSlotNumbersForCarsWithColor       CommandSet = "slot_numbers_for_cars_with_colour"
	CommandSetSlotNumberForRegistrationNumber   CommandSet = "slot_number_for_registration_number"
	CommandSetExit                              CommandSet = "exit"
)
