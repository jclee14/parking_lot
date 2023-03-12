package main

import (
	"errors"
	"strconv"
	"strings"
)

type IParkingLotHandler interface {
	CreateParkingLot(input string) error
	Park(input string) error
	Leave(input string) error
	GetStatus() error
	GetParkedCarNumbersByColor(input string) error
	GetParkedSlotNumbersByColor(input string) error
	GetParkedSlotNumberByCarNumber(input string) error
}

type parkingLotHandler struct {
	parkingLotSvc IParkingLotService
}

func NewParkingLotHandler() parkingLotHandler {
	parkingLotSvc := NewParkingLotService()

	return parkingLotHandler{
		parkingLotSvc: parkingLotSvc,
	}
}

func (h parkingLotHandler) CreateParkingLot(input string) error {
	subStrings := strings.Split(input, " ")
	if len(subStrings) != 2 {
		return errors.New("invalid format")
	}
	slotAmount, err := strconv.Atoi(subStrings[1])
	if err != nil {
		return err
	}
	err = h.parkingLotSvc.CreateParkingLot(slotAmount)
	if err != nil {
		return err
	}

	return nil
}

func (h parkingLotHandler) Park(input string) error {
	subStrings := strings.Split(input, " ")
	if len(subStrings) != 3 {
		return errors.New("invalid format")
	}
	_, err := h.parkingLotSvc.Park(subStrings[1], subStrings[2])
	if err != nil {
		return err
	}

	return nil
}

func (h parkingLotHandler) Leave(input string) error {
	subStrings := strings.Split(input, " ")
	if len(subStrings) != 2 {
		return errors.New("invalid format")
	}
	slotAmount, err := strconv.Atoi(subStrings[1])
	if err != nil {
		return err
	}
	err = h.parkingLotSvc.Leave(slotAmount)
	if err != nil {
		return err
	}

	return nil
}

func (h parkingLotHandler) GetStatus() error {
	err := h.parkingLotSvc.GetStatus()
	if err != nil {
		return err
	}

	return nil
}

func (h parkingLotHandler) GetParkedCarNumbersByColor(input string) error {
	subStrings := strings.Split(input, " ")
	if len(subStrings) != 2 {
		return errors.New("invalid format")
	}
	err := h.parkingLotSvc.GetParkedCarNumbersByColor(subStrings[1])
	if err != nil {
		return err
	}

	return nil
}

func (h parkingLotHandler) GetParkedSlotNumbersByColor(input string) error {
	subStrings := strings.Split(input, " ")
	if len(subStrings) != 2 {
		return errors.New("invalid format")
	}
	err := h.parkingLotSvc.GetParkedSlotNumbersByColor(subStrings[1])
	if err != nil {
		return err
	}

	return nil
}

func (h parkingLotHandler) GetParkedSlotNumberByCarNumber(input string) error {
	subStrings := strings.Split(input, " ")
	if len(subStrings) != 2 {
		return errors.New("invalid format")
	}
	err := h.parkingLotSvc.GetParkedSlotNumberByCarNumber(subStrings[1])
	if err != nil {
		return err
	}

	return nil
}
