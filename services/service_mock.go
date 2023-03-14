package services

import (
	"parking_lot/models"

	"github.com/stretchr/testify/mock"
)

type parkingLotServiceMock struct {
	mock.Mock
}

func NewParkingLotServiceMock() *parkingLotServiceMock {
	return &parkingLotServiceMock{}
}

func (m *parkingLotServiceMock) CreateParkingLot(slotAmount int) ([]*models.ParkingSlot, error) {
	arg := m.Called(slotAmount)
	parkingSlot, ok := arg.Get(0).([]*models.ParkingSlot)
	if !ok {
		parkingSlot = nil
	}
	return parkingSlot, arg.Error(1)
}

func (m *parkingLotServiceMock) Park(carNumber string, carColor string) (int, error) {
	arg := m.Called(carNumber, carColor)
	return arg.Get(0).(int), arg.Error(1)
}

func (m *parkingLotServiceMock) Leave(slotNumber int) error {
	arg := m.Called(slotNumber)
	return arg.Error(0)
}

func (m *parkingLotServiceMock) GetStatus() (string, error) {
	arg := m.Called()
	return arg.Get(0).(string), arg.Error(1)
}

func (m *parkingLotServiceMock) GetParkedCarNumbersByColor(carColor string) (string, error) {
	arg := m.Called(carColor)
	return arg.Get(0).(string), arg.Error(1)
}

func (m *parkingLotServiceMock) GetParkedSlotNumbersByColor(carColor string) (string, error) {
	arg := m.Called(carColor)
	return arg.Get(0).(string), arg.Error(1)
}

func (m *parkingLotServiceMock) GetParkedSlotNumberByCarNumber(carNumber string) (int, error) {
	arg := m.Called(carNumber)
	return arg.Get(0).(int), arg.Error(1)
}

// func (m *parkingLotServiceMock) addCarNumberToColorCache(carNumber string, carColor string) error {
// 	arg := m.Called(carNumber, carColor)
// 	return arg.Error(0)
// }

// func (m *parkingLotServiceMock) removeCarNumberFromColorCache(carNumber string, carColor string) error {
// 	arg := m.Called(carNumber, carColor)
// 	return arg.Error(0)
// }

// func (m *parkingLotServiceMock) addSlotToColorCache(slotNumber int, carColor string) error {
// 	arg := m.Called(slotNumber, carColor)
// 	return arg.Error(0)
// }

// func (m *parkingLotServiceMock) removeSlotFromColorCache(slotNumber int, carColor string) error {
// 	arg := m.Called(slotNumber, carColor)
// 	return arg.Error(0)
// }

// func (m *parkingLotServiceMock) addSlotToCarNumberCache(slotNumber int, carNumber string) error {
// 	arg := m.Called(slotNumber, carNumber)
// 	return arg.Error(0)
// }

// func (m *parkingLotServiceMock) removeSlotFromCarNumberCache(carNumber string) error {
// 	arg := m.Called(carNumber)
// 	return arg.Error(0)
// }
