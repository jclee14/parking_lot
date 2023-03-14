package handlers

import (
	"errors"
	"fmt"
	"parking_lot/models"
	"parking_lot/services"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHandler_CreateParkingLot(t *testing.T) {
	t.Run("create parking lot with 6 slots", func(t *testing.T) {
		slotAmout := 6
		parkingLotService := services.NewParkingLotServiceMock()
		parkingLotService.On("CreateParkingLot", slotAmout).Return(
			[]*models.ParkingSlot{
				{ParkedCar: nil},
				{ParkedCar: nil},
				{ParkedCar: nil},
				{ParkedCar: nil},
				{ParkedCar: nil},
				{ParkedCar: nil},
			}, nil)
		parkingLotHandler := NewParkingLotHandler(parkingLotService)

		err := parkingLotHandler.CreateParkingLot(fmt.Sprintf("create_parking_lot %d", slotAmout))
		assert.Nil(t, err)
	})

	t.Run("create parking lot with empty input", func(t *testing.T) {
		parkingLotService := services.NewParkingLotServiceMock()
		parkingLotHandler := NewParkingLotHandler(parkingLotService)

		err := parkingLotHandler.CreateParkingLot("")
		assert.EqualError(t, err, "invalid format")
	})

	t.Run("create parking lot with incompleted input", func(t *testing.T) {
		parkingLotService := services.NewParkingLotServiceMock()
		parkingLotHandler := NewParkingLotHandler(parkingLotService)

		err := parkingLotHandler.CreateParkingLot("create_parking_lot ")
		assert.EqualError(t, err, "strconv.Atoi: parsing \"\": invalid syntax")
	})

	t.Run("create parking lot with incompleted input", func(t *testing.T) {
		parkingLotService := services.NewParkingLotServiceMock()
		parkingLotHandler := NewParkingLotHandler(parkingLotService)

		err := parkingLotHandler.CreateParkingLot("create_parking_lot s")
		assert.EqualError(t, err, "strconv.Atoi: parsing \"s\": invalid syntax")
	})

	t.Run("create parking lot with 0 slot", func(t *testing.T) {
		parkingLotService := services.NewParkingLotServiceMock()
		parkingLotService.On("CreateParkingLot", 0).Return(nil, errors.New("slot amount must be greater than 0"))
		parkingLotHandler := NewParkingLotHandler(parkingLotService)

		err := parkingLotHandler.CreateParkingLot("create_parking_lot 0")
		assert.EqualError(t, err, "slot amount must be greater than 0")
	})
}

func TestHandler_Park(t *testing.T) {
	t.Run("park car at first slot", func(t *testing.T) {
		parkingLotService := services.NewParkingLotServiceMock()
		parkingLotService.On("Park", "KA-01-HH-7777", "Red").Return(1, nil)
		parkingLotHandler := NewParkingLotHandler(parkingLotService)

		err := parkingLotHandler.Park("park KA-01-HH-7777 Red")
		assert.Nil(t, err)
	})

	t.Run("park car with invalid input", func(t *testing.T) {
		parkingLotService := services.NewParkingLotServiceMock()
		parkingLotHandler := NewParkingLotHandler(parkingLotService)

		err := parkingLotHandler.Park("park KA-01-HH-7777 Red test")
		assert.EqualError(t, err, "invalid format")
	})

	t.Run("park car with empty input", func(t *testing.T) {
		parkingLotService := services.NewParkingLotServiceMock()
		parkingLotHandler := NewParkingLotHandler(parkingLotService)

		err := parkingLotHandler.Park("")
		assert.EqualError(t, err, "invalid format")
	})

	t.Run("park car with incompleted input", func(t *testing.T) {
		parkingLotService := services.NewParkingLotServiceMock()
		parkingLotService.On("Park", "KA-01-HH-7777", "").Return(0, errors.New("car's color is empty"))
		parkingLotHandler := NewParkingLotHandler(parkingLotService)

		err := parkingLotHandler.Park("park KA-01-HH-7777 ")
		assert.EqualError(t, err, "car's color is empty")
	})
}

func TestHandler_Leave(t *testing.T) {
	t.Run("leave car from first slot", func(t *testing.T) {
		parkingLotService := services.NewParkingLotServiceMock()
		parkingLotService.On("Leave", 1).Return(nil)
		parkingLotHandler := NewParkingLotHandler(parkingLotService)

		err := parkingLotHandler.Leave("leave 1")
		assert.Nil(t, err)
	})

	t.Run("leave car with invalid slot number", func(t *testing.T) {
		parkingLotService := services.NewParkingLotServiceMock()
		parkingLotService.On("Leave", 0).Return(errors.New("slot number must be greater than 0"))
		parkingLotHandler := NewParkingLotHandler(parkingLotService)

		err := parkingLotHandler.Leave("Leave 0")
		assert.EqualError(t, err, "slot number must be greater than 0")
	})

	t.Run("leave car with invalid input", func(t *testing.T) {
		parkingLotService := services.NewParkingLotServiceMock()
		parkingLotHandler := NewParkingLotHandler(parkingLotService)

		err := parkingLotHandler.Leave("Leave a")
		assert.EqualError(t, err, "strconv.Atoi: parsing \"a\": invalid syntax")
	})

	t.Run("leave car with empty input", func(t *testing.T) {
		parkingLotService := services.NewParkingLotServiceMock()
		parkingLotHandler := NewParkingLotHandler(parkingLotService)

		err := parkingLotHandler.Leave("")
		assert.EqualError(t, err, "invalid format")
	})
}

func TestHandler_GetStatus(t *testing.T) {}

func TestHandler_GetParkedCarNumbersByColor(t *testing.T) {
	t.Run("get parked car numbers by color", func(t *testing.T) {
		parkingLotService := services.NewParkingLotServiceMock()
		parkingLotService.On("GetParkedCarNumbersByColor", "White").Return("KA-01-HH-1234, KA-01-HH-9999, KA-01-P-333", nil)
		parkingLotHandler := NewParkingLotHandler(parkingLotService)

		err := parkingLotHandler.GetParkedCarNumbersByColor("registration_numbers_for_cars_with_colour White")

		assert.Nil(t, err)
	})

	t.Run("get parked car numbers by empty color", func(t *testing.T) {
		parkingLotService := services.NewParkingLotServiceMock()
		parkingLotService.On("GetParkedCarNumbersByColor", "").Return("", errors.New("car's color is invalid"))
		parkingLotHandler := NewParkingLotHandler(parkingLotService)

		err := parkingLotHandler.GetParkedCarNumbersByColor("registration_numbers_for_cars_with_colour ")

		assert.EqualError(t, err, "car's color is invalid")
	})

	t.Run("get parked car numbers by wrong format input", func(t *testing.T) {
		parkingLotService := services.NewParkingLotServiceMock()
		parkingLotHandler := NewParkingLotHandler(parkingLotService)

		err := parkingLotHandler.GetParkedCarNumbersByColor("registration_numbers_for_cars_with_colour White test")

		assert.EqualError(t, err, "invalid format")
	})
}

func TestHandler_GetParkedSlotNumbersByColor(t *testing.T) {
	t.Run("get parked car numbers by color", func(t *testing.T) {
		parkingLotService := services.NewParkingLotServiceMock()
		parkingLotService.On("GetParkedSlotNumbersByColor", "White").Return("1, 2, 4", nil)
		parkingLotHandler := NewParkingLotHandler(parkingLotService)

		err := parkingLotHandler.GetParkedSlotNumbersByColor("slot_numbers_for_cars_with_colour White")

		assert.Nil(t, err)
	})

	t.Run("get parked car numbers by empty color", func(t *testing.T) {
		parkingLotService := services.NewParkingLotServiceMock()
		parkingLotService.On("GetParkedSlotNumbersByColor", "").Return("", errors.New("car's color is invalid"))
		parkingLotHandler := NewParkingLotHandler(parkingLotService)

		err := parkingLotHandler.GetParkedSlotNumbersByColor("slot_numbers_for_cars_with_colour ")

		assert.EqualError(t, err, "car's color is invalid")
	})

	t.Run("get parked car numbers by wrong format input", func(t *testing.T) {
		parkingLotService := services.NewParkingLotServiceMock()
		parkingLotHandler := NewParkingLotHandler(parkingLotService)

		err := parkingLotHandler.GetParkedSlotNumbersByColor("slot_numbers_for_cars_with_colour White test")

		assert.EqualError(t, err, "invalid format")
	})
}

func TestHandler_GetParkedSlotNumberByCarNumber(t *testing.T) {
	t.Run("get parked car numbers by color", func(t *testing.T) {
		parkingLotService := services.NewParkingLotServiceMock()
		parkingLotService.On("GetParkedSlotNumberByCarNumber", "KA-01-HH-3141").Return(6, nil)
		parkingLotHandler := NewParkingLotHandler(parkingLotService)

		err := parkingLotHandler.GetParkedSlotNumberByCarNumber("slot_number_for_registration_number KA-01-HH-3141")

		assert.Nil(t, err)
	})

	t.Run("get parked car numbers by empty color", func(t *testing.T) {
		parkingLotService := services.NewParkingLotServiceMock()
		parkingLotService.On("GetParkedSlotNumberByCarNumber", "").Return(0, errors.New("car's number is invalid"))
		parkingLotHandler := NewParkingLotHandler(parkingLotService)

		err := parkingLotHandler.GetParkedSlotNumberByCarNumber("slot_number_for_registration_number ")

		assert.EqualError(t, err, "car's number is invalid")
	})

	t.Run("get parked car numbers by wrong format input", func(t *testing.T) {
		parkingLotService := services.NewParkingLotServiceMock()
		parkingLotHandler := NewParkingLotHandler(parkingLotService)

		err := parkingLotHandler.GetParkedSlotNumberByCarNumber("slot_number_for_registration_number KA-01-HH-3141 test")

		assert.EqualError(t, err, "invalid format")
	})
}
