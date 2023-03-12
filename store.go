package main

import "errors"

type IParkingLotStore interface {
	CreateParkingLot(slotAmount int) ([]*parkingSlot, error)
	UpdateParkingLot(slotIndex int, car *car) ([]*parkingSlot, error)
	GetParkingSlots() ([]*parkingSlot, error)
	GetParkingSlotDetail(slotIndex int) (*parkingSlot, error)
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

func (svc *parkingLotStore) UpdateParkingLot(slotIndex int, car *car) ([]*parkingSlot, error) {
	if slotIndex < 0 {
		return nil, errors.New("slot number must be greater than 0")
	}
	if slotIndex > len(svc.parkingSlots)-1 {
		return nil, errors.New("slot amount is not available")
	}

	svc.parkingSlots[slotIndex].parkedCar = car
	return svc.parkingSlots, nil
}

func (svc *parkingLotStore) GetParkingSlots() ([]*parkingSlot, error) {
	if svc.parkingSlots == nil {
		return nil, errors.New("parking lot is not created yet")
	}

	return svc.parkingSlots, nil
}

func (svc *parkingLotStore) GetParkingSlotDetail(slotIndex int) (*parkingSlot, error) {
	if svc.parkingSlots == nil {
		return nil, errors.New("parking lot is not created yet")
	}
	if slotIndex < 0 {
		return nil, errors.New("slot number must be greater than 0")
	}
	if slotIndex > len(svc.parkingSlots)-1 {
		return nil, errors.New("slot amount is not available")
	}

	return svc.parkingSlots[slotIndex], nil
}
