package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type IParkingLotService interface {
	CreateParkingLot(slotAmount int) error
	Park(carNumber string, carColor string) (int, error)
	Leave(slotNumber int) error
	GetStatus() error
	GetParkedCarNumbersByColor(carColor string) error
	GetParkedSlotNumbersByColor(carColor string) error
	GetParkedSlotNumberByCarNumber(carNumber string) error

	addCarNumberToColorCache(carNumber string, carColor string) error
	removeCarNumberFromColorCache(carNumber string, carColor string) error
	addSlotToColorCache(slotNumber int, carColor string) error
	removeSlotFromColorCache(slotNumber int, carColor string) error
	addSlotToCarNumberCache(slotNumber int, carNumber string) error
	removeSlotFromCarNumberCache(carNumber string) error
}

type ParkingLotService struct {
	parkingSlots          []*parkingSlot
	carNumbersByColor     map[string]map[string]struct{}
	slotNumbersByColor    map[string]map[int]struct{}
	slotNumberByCarNumber map[string]int
}

func NewParkingLotService() *ParkingLotService {
	return &ParkingLotService{
		carNumbersByColor:     map[string]map[string]struct{}{},
		slotNumbersByColor:    map[string]map[int]struct{}{},
		slotNumberByCarNumber: map[string]int{},
	}
}

func (svc *ParkingLotService) CreateParkingLot(slotAmount int) error {
	if len(svc.parkingSlots) > 0 {
		return errors.New("parking lot was already created")
	}
	if slotAmount < 1 {
		return errors.New("slot amount must be greater than 0")
	}

	for i := 0; i < slotAmount; i++ {
		svc.parkingSlots = append(svc.parkingSlots, &parkingSlot{})
	}
	fmt.Printf("Created a parking lot with %d slots\n", slotAmount)

	return nil
}

func (svc *ParkingLotService) Park(carNumber string, carColor string) (int, error) {
	if len(carNumber) == 0 {
		return 0, errors.New("car's number is empty")
	}
	if len(carColor) == 0 {
		return 0, errors.New("car's color is empty")
	}

	carColor = strings.ToLower(carColor)

	for idx, slot := range svc.parkingSlots {
		if slot.parkedCar != nil {
			continue
		}

		slotNumber := idx + 1

		slot.parkedCar = &car{}
		slot.parkedCar.registrationNumber = carNumber
		slot.parkedCar.color = carColor

		err := svc.addCarNumberToColorCache(carNumber, carColor)
		if err != nil {
			slot.parkedCar = nil
			return 0, errors.Join(err)
		}
		err = svc.addSlotToColorCache(slotNumber, carColor)
		if err != nil {
			slot.parkedCar = nil
			return 0, errors.Join(err)
		}
		err = svc.addSlotToCarNumberCache(slotNumber, carNumber)
		if err != nil {
			slot.parkedCar = nil
			return 0, errors.Join(err)
		}

		fmt.Printf("Allocated slot number: %d\n", slotNumber)
		return slotNumber, nil
	}

	return 0, errors.New("parking lot is full")
}

func (svc *ParkingLotService) Leave(slotNumber int) error {
	if slotNumber < 1 || slotNumber > len(svc.parkingSlots) {
		return errors.New("slot number is invalid")
	}

	carColor := svc.parkingSlots[slotNumber-1].parkedCar.color
	carNumber := svc.parkingSlots[slotNumber-1].parkedCar.registrationNumber
	err := svc.removeCarNumberFromColorCache(carNumber, carColor)
	if err != nil {
		return errors.Join(err)
	}
	err = svc.removeSlotFromColorCache(slotNumber, carColor)
	if err != nil {
		return errors.Join(err)
	}
	err = svc.removeSlotFromCarNumberCache(svc.parkingSlots[slotNumber-1].parkedCar.registrationNumber)
	if err != nil {
		return errors.Join(err)
	}

	svc.parkingSlots[slotNumber-1].parkedCar = nil
	fmt.Printf("Slot number %d is free\n", slotNumber)

	return nil
}

func (svc *ParkingLotService) GetStatus() error {
	fmt.Println("Slot No.\tRegistration No\t\tColour")
	for idx, slot := range svc.parkingSlots {
		if slot.parkedCar == nil {
			continue
		}

		fmt.Printf("%d\t\t%s\t\t%s\n", idx, slot.parkedCar.registrationNumber, slot.parkedCar.color)
	}

	return nil
}

func (svc *ParkingLotService) GetParkedCarNumbersByColor(carColor string) error {
	if len(carColor) == 0 {
		return errors.New("car's color is invalid")
	}

	cData, ok := svc.carNumbersByColor[carColor]
	if !ok {
		return errors.New("not found")
	}

	results := make([]string, 0, len(cData))
	for carNumber := range cData {
		if len(carNumber) == 0 {
			continue
		}
		results = append(results, carNumber)
	}
	fmt.Println(strings.Join(results, ", "))

	return nil
}

func (svc *ParkingLotService) GetParkedSlotNumbersByColor(carColor string) error {
	if len(carColor) == 0 {
		return errors.New("car's color is invalid")
	}

	cData, ok := svc.slotNumbersByColor[carColor]
	if !ok {
		return errors.New("not found")
	}

	results := make([]string, 0, len(cData))
	for slotNumber := range cData {
		slotStr := strconv.Itoa(slotNumber)
		if len(slotStr) == 0 {
			continue
		}
		results = append(results, strconv.Itoa(slotNumber))
	}
	fmt.Println(strings.Join(results, ", "))

	return nil
}

func (svc *ParkingLotService) GetParkedSlotNumberByCarNumber(carNumber string) error {
	if len(carNumber) == 0 {
		return errors.New("car's number is invalid")
	}

	slotNumber, ok := svc.slotNumberByCarNumber[carNumber]
	if !ok {
		return errors.New("not found")
	}

	fmt.Println(slotNumber)

	return nil
}

func (svc *ParkingLotService) addCarNumberToColorCache(carNumber string, carColor string) error {
	if len(carNumber) == 0 {
		return errors.New("car's number is empty")
	}
	if len(carColor) == 0 {
		return errors.New("car's color is empty")
	}

	cData, ok := svc.carNumbersByColor[carColor]
	if !ok {
		svc.carNumbersByColor[carColor] = map[string]struct{}{
			carNumber: {},
		}
	} else {
		cData[carNumber] = struct{}{}
		svc.carNumbersByColor[carColor] = cData
	}

	return nil
}

func (svc *ParkingLotService) removeCarNumberFromColorCache(carNumber string, carColor string) error {
	if len(carNumber) == 0 {
		return errors.New("car's number is empty")
	}
	if len(carColor) == 0 {
		return errors.New("car's color is empty")
	}

	if cData, ok := svc.carNumbersByColor[carColor]; ok {
		delete(cData, carNumber)
		svc.carNumbersByColor[carColor] = cData
	}

	return nil
}

func (svc *ParkingLotService) addSlotToColorCache(slotNumber int, carColor string) error {
	if slotNumber < 1 || slotNumber > len(svc.parkingSlots) {
		return errors.New("slot number is invalid")
	}
	if len(carColor) == 0 {
		return errors.New("car's color is empty")
	}

	cData, ok := svc.slotNumbersByColor[carColor]
	if !ok {
		svc.slotNumbersByColor[carColor] = map[int]struct{}{
			slotNumber: {},
		}
	} else {
		cData[slotNumber] = struct{}{}
		svc.slotNumbersByColor[carColor] = cData
	}

	return nil
}

func (svc *ParkingLotService) removeSlotFromColorCache(slotNumber int, carColor string) error {
	if slotNumber < 1 || slotNumber > len(svc.parkingSlots) {
		return errors.New("slot number is invalid")
	}
	if len(carColor) == 0 {
		return errors.New("car's color is empty")
	}

	if cData, ok := svc.slotNumbersByColor[carColor]; ok {
		delete(cData, slotNumber)
		svc.slotNumbersByColor[carColor] = cData
	}

	return nil
}

func (svc *ParkingLotService) addSlotToCarNumberCache(slotNumber int, carNumber string) error {
	if slotNumber < 1 || slotNumber > len(svc.parkingSlots) {
		return errors.New("slot number is invalid")
	}
	if len(carNumber) == 0 {
		return errors.New("car's number is empty")
	}

	svc.slotNumberByCarNumber[carNumber] = slotNumber
	return nil
}

func (svc *ParkingLotService) removeSlotFromCarNumberCache(carNumber string) error {
	if len(carNumber) == 0 {
		return errors.New("car's number is empty")
	}

	delete(svc.slotNumberByCarNumber, carNumber)
	return nil
}
