package stores

import (
	"parking_lot/models"

	"github.com/stretchr/testify/mock"
)

type parkingLotStoreMock struct {
	mock.Mock
}

func NewParkingLotStoreMock() *parkingLotStoreMock {
	return &parkingLotStoreMock{}
}

func (m *parkingLotStoreMock) CreateParkingLot(slotAmount int) ([]*models.ParkingSlot, error) {
	arg := m.Called(slotAmount)
	parkingSlot, ok := arg.Get(0).([]*models.ParkingSlot)
	if !ok {
		parkingSlot = nil
	}
	return parkingSlot, arg.Error(1)
}

func (m *parkingLotStoreMock) UpdateParkingLot(slotNumber int, car *models.Car) ([]*models.ParkingSlot, error) {
	arg := m.Called(slotNumber, car)
	parkingSlot, ok := arg.Get(0).([]*models.ParkingSlot)
	if !ok {
		parkingSlot = nil
	}
	return parkingSlot, arg.Error(1)
}

func (m *parkingLotStoreMock) GetParkingSlots() ([]*models.ParkingSlot, error) {
	arg := m.Called()
	parkingSlot, ok := arg.Get(0).([]*models.ParkingSlot)
	if !ok {
		parkingSlot = nil
	}
	return parkingSlot, arg.Error(1)
}

func (m *parkingLotStoreMock) GetParkingSlotDetail(slotNumber int) (*models.ParkingSlot, error) {
	arg := m.Called(slotNumber)
	parkingSlot, ok := arg.Get(0).(*models.ParkingSlot)
	if !ok {
		parkingSlot = nil
	}
	return parkingSlot, arg.Error(1)
}
