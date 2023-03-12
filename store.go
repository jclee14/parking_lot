package main

import "errors"

type IParkingLotStore interface {
	CreateParkingLot(slotAmount int) ([]*parkingSlot, error)
	UpdateParkingLot(slotNumber int, car *car) ([]*parkingSlot, error)
	GetParkingSlots() ([]*parkingSlot, error)
	GetParkingSlotDetail(slotNumber int) (*parkingSlot, error)
}

type parkingLotStore struct {
	parkingSlots []*parkingSlot
}

func NewParkingLotStore() *parkingLotStore {
	return &parkingLotStore{}
}

func (svc *parkingLotStore) CreateParkingLot(slotAmount int) ([]*parkingSlot, error) {
	if slotAmount < 1 {
		return nil, errors.New("slot amount must be greater than 0")
	}

	parkingSlots, _ := svc.GetParkingSlots()
	if len(parkingSlots) > 0 {
		return nil, errors.New("parking lot was already created")
	}

	for i := 0; i < slotAmount; i++ {
		svc.parkingSlots = append(svc.parkingSlots, &parkingSlot{})
	}

	return svc.parkingSlots, nil
}

func (svc *parkingLotStore) UpdateParkingLot(slotNumber int, car *car) ([]*parkingSlot, error) {
	if slotNumber < 0 {
		return nil, errors.New("slot number must be greater than 0")
	}
	if slotNumber > len(svc.parkingSlots)-1 {
		return nil, errors.New("slot amount is not available")
	}

	svc.parkingSlots[slotNumber].parkedCar = car
	return svc.parkingSlots, nil
}

func (svc *parkingLotStore) GetParkingSlots() ([]*parkingSlot, error) {
	if svc.parkingSlots == nil {
		return nil, errors.New("parking lot is not created yet")
	}

	return svc.parkingSlots, nil
}

func (svc *parkingLotStore) GetParkingSlotDetail(slotNumber int) (*parkingSlot, error) {
	if svc.parkingSlots == nil {
		return nil, errors.New("parking lot is not created yet")
	}
	if slotNumber < 0 {
		return nil, errors.New("slot number must be greater than 0")
	}
	if slotNumber > len(svc.parkingSlots)-1 {
		return nil, errors.New("slot amount is not available")
	}

	return svc.parkingSlots[slotNumber], nil
}
