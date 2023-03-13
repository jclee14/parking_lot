package main

import "github.com/stretchr/testify/mock"

type parkingLotStoreMock struct {
	mock.Mock
}

func NewParkingLotStoreMock() *parkingLotStoreMock {
	return &parkingLotStoreMock{}
}

func (m *parkingLotStoreMock) CreateParkingLot(slotAmount int) ([]*parkingSlot, error) {
	arg := m.Called(slotAmount)
	parkingSlot, ok := arg.Get(0).([]*parkingSlot)
	if !ok {
		parkingSlot = nil
	}
	return parkingSlot, arg.Error(1)
}

func (m *parkingLotStoreMock) UpdateParkingLot(slotNumber int, car *car) ([]*parkingSlot, error) {
	arg := m.Called(slotNumber, car)
	parkingSlot, ok := arg.Get(0).([]*parkingSlot)
	if !ok {
		parkingSlot = nil
	}
	return parkingSlot, arg.Error(1)
}

func (m *parkingLotStoreMock) GetParkingSlots() ([]*parkingSlot, error) {
	arg := m.Called()
	parkingSlot, ok := arg.Get(0).([]*parkingSlot)
	if !ok {
		parkingSlot = nil
	}
	return parkingSlot, arg.Error(1)
}

func (m *parkingLotStoreMock) GetParkingSlotDetail(slotNumber int) (*parkingSlot, error) {
	arg := m.Called(slotNumber)
	parkingSlot, ok := arg.Get(0).(*parkingSlot)
	if !ok {
		parkingSlot = nil
	}
	return parkingSlot, arg.Error(1)
}
