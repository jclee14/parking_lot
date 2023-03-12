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
	return arg.Get(0).([]*parkingSlot), arg.Error(1)
}

func (m *parkingLotStoreMock) UpdateParkingLot(slotNumber int, car *car) ([]*parkingSlot, error) {
	arg := m.Called(slotNumber, car)
	return arg.Get(0).([]*parkingSlot), arg.Error(1)
}

func (m *parkingLotStoreMock) GetParkingSlots() ([]*parkingSlot, error) {
	arg := m.Called()
	return arg.Get(0).([]*parkingSlot), arg.Error(1)
}

func (m *parkingLotStoreMock) GetParkingSlotDetail(slotNumber int) (*parkingSlot, error) {
	arg := m.Called(slotNumber)
	return arg.Get(0).(*parkingSlot), arg.Error(1)
}
