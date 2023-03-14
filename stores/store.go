package stores

import (
	"errors"
	"parking_lot/models"
)

type IParkingLotStore interface {
	CreateParkingLot(slotAmount int) ([]*models.ParkingSlot, error)
	UpdateParkingLot(slotIndex int, car *models.Car) ([]*models.ParkingSlot, error)
	GetParkingSlots() ([]*models.ParkingSlot, error)
	GetParkingSlotDetail(slotIndex int) (*models.ParkingSlot, error)
}

type parkingLotStore struct {
	parkingSlots []*models.ParkingSlot
}

func NewParkingLotStore() *parkingLotStore {
	return &parkingLotStore{}
}

func (svc *parkingLotStore) CreateParkingLot(slotAmount int) ([]*models.ParkingSlot, error) {
	if slotAmount < 1 {
		return nil, errors.New("slot amount must be greater than 0")
	}

	parkingSlots, _ := svc.GetParkingSlots()
	if len(parkingSlots) > 0 {
		return nil, errors.New("parking lot was already created")
	}

	for i := 0; i < slotAmount; i++ {
		svc.parkingSlots = append(svc.parkingSlots, &models.ParkingSlot{})
	}

	return svc.parkingSlots, nil
}

func (svc *parkingLotStore) UpdateParkingLot(slotIndex int, car *models.Car) ([]*models.ParkingSlot, error) {
	if slotIndex < 0 {
		return nil, errors.New("slot number must be greater than 0")
	}
	if slotIndex > len(svc.parkingSlots)-1 {
		return nil, errors.New("slot amount is not available")
	}

	svc.parkingSlots[slotIndex].ParkedCar = car
	return svc.parkingSlots, nil
}

func (svc *parkingLotStore) GetParkingSlots() ([]*models.ParkingSlot, error) {
	if svc.parkingSlots == nil {
		return nil, errors.New("parking lot is not created yet")
	}

	return svc.parkingSlots, nil
}

func (svc *parkingLotStore) GetParkingSlotDetail(slotIndex int) (*models.ParkingSlot, error) {
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
