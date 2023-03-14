package handlers

import (
	"errors"
	"fmt"
	"parking_lot/services"
	"strconv"
	"strings"
)

// type IParkingLotHandler interface {
// 	CreateParkingLot(input string) error
// 	Park(input string) error
// 	Leave(input string) error
// 	GetStatus() error
// 	GetParkedCarNumbersByColor(input string) error
// 	GetParkedSlotNumbersByColor(input string) error
// 	GetParkedSlotNumberByCarNumber(input string) error
// }

type ParkingLotHandler struct {
	parkingLotSvc services.IParkingLotService
}

func NewParkingLotHandler(parkingLotSvc services.IParkingLotService) ParkingLotHandler {
	return ParkingLotHandler{
		parkingLotSvc: parkingLotSvc,
	}
}

func (h ParkingLotHandler) CreateParkingLot(input string) error {
	subStrings := strings.Split(input, " ")
	if len(subStrings) != 2 {
		return errors.New("invalid format")
	}
	slotAmount, err := strconv.Atoi(subStrings[1])
	if err != nil {
		return err
	}
	parkingSlots, err := h.parkingLotSvc.CreateParkingLot(slotAmount)
	if err != nil {
		return err
	}

	fmt.Printf("Created a parking lot with %d slots\n", len(parkingSlots))
	return nil
}

func (h ParkingLotHandler) Park(input string) error {
	subStrings := strings.Split(input, " ")
	if len(subStrings) != 3 {
		return errors.New("invalid format")
	}
	parkedSlotNumber, err := h.parkingLotSvc.Park(subStrings[1], subStrings[2])
	if err != nil {
		return err
	}

	fmt.Printf("Allocated slot number: %d\n", parkedSlotNumber)
	return nil
}

func (h ParkingLotHandler) Leave(input string) error {
	subStrings := strings.Split(input, " ")
	if len(subStrings) != 2 {
		return errors.New("invalid format")
	}
	slotNumber, err := strconv.Atoi(subStrings[1])
	if err != nil {
		return err
	}
	err = h.parkingLotSvc.Leave(slotNumber)
	if err != nil {
		return err
	}

	fmt.Printf("Slot number %d is free\n", slotNumber)
	return nil
}

func (h ParkingLotHandler) GetStatus() error {
	status, err := h.parkingLotSvc.GetStatus()
	if err != nil {
		return err
	}

	fmt.Println("Slot No.\tRegistration No\t\tColour")
	fmt.Println(status)

	return nil
}

func (h ParkingLotHandler) GetParkedCarNumbersByColor(input string) error {
	subStrings := strings.Split(input, " ")
	if len(subStrings) != 2 {
		return errors.New("invalid format")
	}
	data, err := h.parkingLotSvc.GetParkedCarNumbersByColor(subStrings[1])
	if err != nil {
		return err
	}

	fmt.Println(data)

	return nil
}

func (h ParkingLotHandler) GetParkedSlotNumbersByColor(input string) error {
	subStrings := strings.Split(input, " ")
	if len(subStrings) != 2 {
		return errors.New("invalid format")
	}
	data, err := h.parkingLotSvc.GetParkedSlotNumbersByColor(subStrings[1])
	if err != nil {
		return err
	}

	fmt.Println(data)

	return nil
}

func (h ParkingLotHandler) GetParkedSlotNumberByCarNumber(input string) error {
	subStrings := strings.Split(input, " ")
	if len(subStrings) != 2 {
		return errors.New("invalid format")
	}
	data, err := h.parkingLotSvc.GetParkedSlotNumberByCarNumber(subStrings[1])
	if err != nil {
		return err
	}

	fmt.Println(data)

	return nil
}
