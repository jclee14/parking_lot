package main

import (
	"errors"
	"fmt"
)

type IParkingLotService interface {
	Park(carNumber string, carColor string) (int, error)
	Leave(slotNumber int) error
	GetStatus() error
	addSlotToColorCache(slotNumber int, carColor string) error
	removeSlotFromColorCache(slotNumber int, carColor string) error
	addSlotToCarNumberCache(slotNumber int, carNumber string) error
	removeSlotFromCarNumberCache(carNumber string) error
}

type ParkingLotService struct {
	parkingSlots          []*parkingSlot
	slotNumbersByColor    map[string]map[int]struct{}
	slotNumberByCarNumber map[string]int
}

func NewParkingLotService(slotAmount int) (*ParkingLotService, error) {
	if slotAmount < 1 {
		return nil, errors.New("slot amount must be greater than 0")
	}

	return &ParkingLotService{
		parkingSlots:          make([]*parkingSlot, slotAmount),
		slotNumbersByColor:    map[string]map[int]struct{}{},
		slotNumberByCarNumber: map[string]int{},
	}, nil
}

func (svc *ParkingLotService) Park(carNumber string, carColor string) (int, error) {
	if len(carNumber) == 0 {
		return 0, errors.New("car's number is empty")
	}
	if len(carColor) == 0 {
		return 0, errors.New("car's color is empty")
	}

	for idx, slot := range svc.parkingSlots {
		if slot.parkedCar != nil {
			continue
		}

		slot.parkedCar.registrationNumber = carNumber
		slot.parkedCar.color = carColor

		err := svc.addSlotToColorCache(idx, carColor)
		if err != nil {
			slot.parkedCar = nil
			return 0, errors.Join(err)
		}
		err = svc.addSlotToCarNumberCache(idx, carNumber)
		if err != nil {
			slot.parkedCar = nil
			return 0, errors.Join(err)
		}

		return idx, nil
	}

	return 0, errors.New("parking lot is full")
}

func (svc *ParkingLotService) Leave(slotNumber int) error {
	if slotNumber < 1 || slotNumber > len(svc.parkingSlots) {
		return errors.New("slot number is invalid")
	}

	err := svc.removeSlotFromColorCache(slotNumber, svc.parkingSlots[slotNumber-1].parkedCar.color)
	if err != nil {
		return errors.Join(err)
	}
	err = svc.removeSlotFromCarNumberCache(svc.parkingSlots[slotNumber-1].parkedCar.registrationNumber)
	if err != nil {
		return errors.Join(err)
	}

	svc.parkingSlots[slotNumber-1].parkedCar = nil
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
