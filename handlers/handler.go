package handlers

import (
	"fmt"
	"parking_lot/services"
	"strconv"
)

type ParkingLotHandler struct {
	parkingLotSvc services.IParkingLotService
}

func NewParkingLotHandler(parkingLotSvc services.IParkingLotService) ParkingLotHandler {
	return ParkingLotHandler{
		parkingLotSvc: parkingLotSvc,
	}
}

func (h ParkingLotHandler) CreateParkingLot(input string) error {
	slotAmount, err := strconv.Atoi(input)
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

func (h ParkingLotHandler) Park(carNumber string, carColor string) error {
	parkedSlotNumber, err := h.parkingLotSvc.Park(carNumber, carColor)
	if err != nil {
		return err
	}

	fmt.Printf("Allocated slot number: %d\n", parkedSlotNumber)
	return nil
}

func (h ParkingLotHandler) Leave(input string) error {
	slotNumber, err := strconv.Atoi(input)
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

func (h ParkingLotHandler) GetParkedCarNumbersByColor(CarColor string) error {
	data, err := h.parkingLotSvc.GetParkedCarNumbersByColor(CarColor)
	if err != nil {
		return err
	}

	fmt.Println(data)

	return nil
}

func (h ParkingLotHandler) GetParkedSlotNumbersByColor(CarColor string) error {
	data, err := h.parkingLotSvc.GetParkedSlotNumbersByColor(CarColor)
	if err != nil {
		return err
	}

	fmt.Println(data)

	return nil
}

func (h ParkingLotHandler) GetParkedSlotNumberByCarNumber(carNumber string) error {
	data, err := h.parkingLotSvc.GetParkedSlotNumberByCarNumber(carNumber)
	if err != nil {
		return err
	}

	fmt.Println(data)

	return nil
}
