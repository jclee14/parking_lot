package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

func main() {
	parkingLotStore := NewParkingLotStore()
	parkingLotSvc := NewParkingLotService(parkingLotStore)
	parkingLotHandler := NewParkingLotHandler(parkingLotSvc)
	start(parkingLotHandler)
}

func start(parkingLotHandler IParkingLotHandler) {
	reader := bufio.NewReader(os.Stdin)

	for {
		input, _ := reader.ReadString('\n')
		// convert CRLF to LF
		// input = strings.Replace(input, "\n", "", -1)
		input = strings.TrimSpace(strings.Replace(input, "\r\n", "", -1))
		if strings.HasPrefix(input, "create_parking_lot") {
			err := parkingLotHandler.CreateParkingLot(input)
			if err != nil {
				fmt.Println(err)
				continue
			}
		} else if strings.HasPrefix(input, "park") {
			err := parkingLotHandler.Park(input)
			if err != nil {
				fmt.Println(err)
				continue
			}
		} else if strings.HasPrefix(input, "leave") {
			err := parkingLotHandler.Leave(input)
			if err != nil {
				fmt.Println(err)
				continue
			}
		} else if strings.HasPrefix(input, "status") {
			err := parkingLotHandler.GetStatus()
			if err != nil {
				fmt.Println(err)
				continue
			}
		} else if strings.HasPrefix(input, "registration_numbers_for_cars_with_colour") {
			err := parkingLotHandler.GetParkedCarNumbersByColor(input)
			if err != nil {
				fmt.Println(err)
				continue
			}
		} else if strings.HasPrefix(input, "slot_numbers_for_cars_with_colour") {
			err := parkingLotHandler.GetParkedSlotNumbersByColor(input)
			if err != nil {
				fmt.Println(err)
				continue
			}
		} else if strings.HasPrefix(input, "slot_number_for_registration_number") {
			err := parkingLotHandler.GetParkedSlotNumberByCarNumber(input)
			if err != nil {
				fmt.Println(err)
				continue
			}
		} else if strings.Contains(input, "exit") {
			break
		} else {
			fmt.Println(errors.New("command not found"))
			continue
		}
	}
}
