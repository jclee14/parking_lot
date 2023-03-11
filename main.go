package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var parkingLotSvc IParkingLotService = NewParkingLotService()
	start(parkingLotSvc)
}

func start(parkingLotSvc IParkingLotService) {
	reader := bufio.NewReader(os.Stdin)

	for {
		text, _ := reader.ReadString('\n')
		// convert CRLF to LF
		// text = strings.Replace(text, "\n", "", -1)
		text = strings.TrimSpace(strings.Replace(text, "\r\n", "", -1))
		if strings.HasPrefix(text, "create_parking_lot") {
			subStrings := strings.Split(text, " ")
			if len(subStrings) != 2 {
				fmt.Println(errors.New("invalid format"))
				continue
			}
			slotAmount, err := strconv.Atoi(subStrings[1])
			if err != nil {
				fmt.Println(err)
				continue
			}
			err = parkingLotSvc.CreateParkingLot(slotAmount)
			if err != nil {
				fmt.Println(err)
				continue
			}
		} else if strings.HasPrefix(text, "park") {
			subStrings := strings.Split(text, " ")
			if len(subStrings) != 3 {
				fmt.Println(errors.New("invalid format"))
				continue
			}
			_, err := parkingLotSvc.Park(subStrings[1], subStrings[2])
			if err != nil {
				fmt.Println(err)
				continue
			}
		} else if strings.HasPrefix(text, "leave") {
			subStrings := strings.Split(text, " ")
			if len(subStrings) != 2 {
				fmt.Println(errors.New("invalid format"))
				continue
			}
			slotAmount, err := strconv.Atoi(subStrings[1])
			if err != nil {
				fmt.Println(err)
				continue
			}
			err = parkingLotSvc.Leave(slotAmount)
			if err != nil {
				fmt.Println(err)
				continue
			}
		} else if strings.HasPrefix(text, "status") {
			err := parkingLotSvc.PrintStatus()
			if err != nil {
				fmt.Println(err)
				continue
			}
		} else if strings.HasPrefix(text, "registration_numbers_for_cars_with_colour") {
			subStrings := strings.Split(text, " ")
			if len(subStrings) != 2 {
				fmt.Println(errors.New("invalid format"))
				continue
			}
			err := parkingLotSvc.PrintParkedCarNumbersByColor(subStrings[1])
			if err != nil {
				fmt.Println(err)
				continue
			}
		} else if strings.HasPrefix(text, "slot_numbers_for_cars_with_colour") {
			subStrings := strings.Split(text, " ")
			if len(subStrings) != 2 {
				fmt.Println(errors.New("invalid format"))
				continue
			}
			err := parkingLotSvc.PrintParkedSlotNumbersByColor(subStrings[1])
			if err != nil {
				fmt.Println(err)
				continue
			}
		} else if strings.HasPrefix(text, "slot_number_for_registration_number") {
			subStrings := strings.Split(text, " ")
			if len(subStrings) != 2 {
				fmt.Println(errors.New("invalid format"))
				continue
			}
			err := parkingLotSvc.PrintParkedSlotNumberByCarNumber(subStrings[1])
			if err != nil {
				fmt.Println(err)
				continue
			}
		} else if strings.Contains(text, "exit") {
			break
		} else {
			fmt.Println(errors.New("command not found"))
			continue
		}
	}
}
