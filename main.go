package main

import (
	"errors"
	"fmt"
	"parking_lot/handlers"
	"parking_lot/models"
	"parking_lot/services"
	"parking_lot/stores"
)

func main() {
	parkingLotStore := stores.NewParkingLotStore()
	parkingLotSvc := services.NewParkingLotService(parkingLotStore)
	parkingLotHandler := handlers.NewParkingLotHandler(parkingLotSvc)
	start(parkingLotHandler)
}

func start(parkingLotHandler handlers.ParkingLotHandler) {
	for {
		var command models.CommandSet
		var arg1, arg2 string
		fmt.Scanf("%s %s %s\n", &command, &arg1, &arg2)
		exitProgram := false

		switch command {
		case models.CommandSetCreateParkingLot:
			err := parkingLotHandler.CreateParkingLot(arg1)
			if err != nil {
				fmt.Println(err)
			}
		case models.CommandSetPark:
			err := parkingLotHandler.Park(arg1, arg2)
			if err != nil {
				fmt.Println(err)
			}
		case models.CommandSetLeave:
			err := parkingLotHandler.Leave(arg1)
			if err != nil {
				fmt.Println(err)
			}
		case models.CommandSetStatus:
			err := parkingLotHandler.GetStatus()
			if err != nil {
				fmt.Println(err)
			}
		case models.CommandSetRegistrationNumberForCarWithColor:
			err := parkingLotHandler.GetParkedCarNumbersByColor(arg1)
			if err != nil {
				fmt.Println(err)
			}
		case models.CommandSetSlotNumbersForCarsWithColor:
			err := parkingLotHandler.GetParkedSlotNumbersByColor(arg1)
			if err != nil {
				fmt.Println(err)
			}
		case models.CommandSetSlotNumberForRegistrationNumber:
			err := parkingLotHandler.GetParkedSlotNumberByCarNumber(arg1)
			if err != nil {
				fmt.Println(err)
			}
		case models.CommandSetExit:
			exitProgram = true
		default:
			fmt.Println(errors.New("command not found"))
		}

		if exitProgram {
			break
		}
	}
}
