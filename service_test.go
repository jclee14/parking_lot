package main

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestService_CreateParkingLot(t *testing.T) {
	t.Run("create parking lot with 6 slots", func(t *testing.T) {
		parkingLotStore := NewParkingLotStoreMock()
		parkingLotStore.On("CreateParkingLot", 6).Return(
			[]*parkingSlot{
				{parkedCar: nil},
				{parkedCar: nil},
				{parkedCar: nil},
				{parkedCar: nil},
				{parkedCar: nil},
				{parkedCar: nil},
			}, nil)
		parkingLotService := NewParkingLotService(parkingLotStore)

		parkingSlots, err := parkingLotService.CreateParkingLot(6)
		expected := []*parkingSlot{
			{parkedCar: nil},
			{parkedCar: nil},
			{parkedCar: nil},
			{parkedCar: nil},
			{parkedCar: nil},
			{parkedCar: nil},
		}

		assert.Equal(t, expected, parkingSlots)
		assert.Nil(t, err)
	})

	t.Run("create parking lot with 0 slot", func(t *testing.T) {
		parkingLotStore := NewParkingLotStoreMock()
		parkingLotStore.On("CreateParkingLot", 0).Return([]*parkingSlot{}, errors.New("slot amount must be greater than 0"))
		parkingLotService := NewParkingLotService(parkingLotStore)

		parkingSlots, err := parkingLotService.CreateParkingLot(0)

		assert.Nil(t, parkingSlots)
		assert.Error(t, err)
	})
}

func TestService_Park(t *testing.T) {
	parkingLotStore := NewParkingLotStoreMock()
	parkingLotStore.On("GetParkingSlots").Return(
		[]*parkingSlot{
			{parkedCar: nil},
			{parkedCar: nil},
			{parkedCar: nil},
			{parkedCar: nil},
			{parkedCar: nil},
			{parkedCar: nil},
		}, nil)
	parkingLotStore.On("UpdateParkingLot", 0, &car{registrationNumber: "KA-01-HH-2701", color: "blue"}).Return(
		[]*parkingSlot{
			{parkedCar: &car{registrationNumber: "KA-01-HH-2701", color: "blue"}},
			{parkedCar: nil},
			{parkedCar: nil},
			{parkedCar: nil},
			{parkedCar: nil},
			{parkedCar: nil},
		}, nil)
	parkingLotService := NewParkingLotService(parkingLotStore)

	parkedSlotNumber, err := parkingLotService.Park("KA-01-HH-2701", "blue")
	expected := 1

	assert.Equal(t, expected, parkedSlotNumber)
	assert.Nil(t, err)
}
