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

		err := parkingLotHandler.CreateParkingLot(fmt.Sprintf("%d", slotAmout))
		assert.Nil(t, err)
	})

	t.Run("create parking lot with empty input", func(t *testing.T) {
		parkingLotService := services.NewParkingLotServiceMock()
		parkingLotHandler := NewParkingLotHandler(parkingLotService)

		err := parkingLotHandler.CreateParkingLot("")
		assert.EqualError(t, err, "strconv.Atoi: parsing \"\": invalid syntax")
	})

	t.Run("create parking lot with 0 slot", func(t *testing.T) {
		slotAmout := 0
		parkingLotService := services.NewParkingLotServiceMock()
		parkingLotService.On("CreateParkingLot", slotAmout).Return(nil, errors.New("slot amount must be greater than 0"))
		parkingLotHandler := NewParkingLotHandler(parkingLotService)

		err := parkingLotHandler.CreateParkingLot(fmt.Sprintf("%d", slotAmout))
		assert.EqualError(t, err, "slot amount must be greater than 0")
	})
}

func TestHandler_Park(t *testing.T) {
	t.Run("park car at first slot", func(t *testing.T) {
		parkingLotService := services.NewParkingLotServiceMock()
		parkingLotService.On("Park", "KA-01-HH-7777", "Red").Return(1, nil)
		parkingLotHandler := NewParkingLotHandler(parkingLotService)

		err := parkingLotHandler.Park("KA-01-HH-7777", "Red")
		assert.Nil(t, err)
	})

	t.Run("park car with empty car's color", func(t *testing.T) {
		parkingLotService := services.NewParkingLotServiceMock()
		parkingLotService.On("Park", "KA-01-HH-7777", "").Return(0, errors.New("car's color is empty"))
		parkingLotHandler := NewParkingLotHandler(parkingLotService)

		err := parkingLotHandler.Park("KA-01-HH-7777", "")
		assert.EqualError(t, err, "car's color is empty")
	})

	t.Run("park car with empty car's number", func(t *testing.T) {
		parkingLotService := services.NewParkingLotServiceMock()
		parkingLotService.On("Park", "", "Red").Return(0, errors.New("car's number is empty"))
		parkingLotHandler := NewParkingLotHandler(parkingLotService)

		err := parkingLotHandler.Park("", "Red")
		assert.EqualError(t, err, "car's number is empty")
	})
}

func TestHandler_Leave(t *testing.T) {
	t.Run("leave car from first slot", func(t *testing.T) {
		slotNumber := 1
		parkingLotService := services.NewParkingLotServiceMock()
		parkingLotService.On("Leave", slotNumber).Return(nil)
		parkingLotHandler := NewParkingLotHandler(parkingLotService)

		err := parkingLotHandler.Leave(fmt.Sprintf("%d", slotNumber))
		assert.Nil(t, err)
	})

	t.Run("leave car with invalid slot number", func(t *testing.T) {
		slotNumber := 0
		parkingLotService := services.NewParkingLotServiceMock()
		parkingLotService.On("Leave", slotNumber).Return(errors.New("slot number must be greater than 0"))
		parkingLotHandler := NewParkingLotHandler(parkingLotService)

		err := parkingLotHandler.Leave(fmt.Sprintf("%d", slotNumber))
		assert.EqualError(t, err, "slot number must be greater than 0")
	})

	t.Run("leave car with empty input", func(t *testing.T) {
		parkingLotService := services.NewParkingLotServiceMock()
		parkingLotHandler := NewParkingLotHandler(parkingLotService)

		err := parkingLotHandler.Leave("")
		assert.EqualError(t, err, "strconv.Atoi: parsing \"\": invalid syntax")
	})
}

func TestHandler_GetStatus(t *testing.T) {}

func TestHandler_GetParkedCarNumbersByColor(t *testing.T) {
	t.Run("get parked car numbers by car's color", func(t *testing.T) {
		carColor := "White"
		parkingLotService := services.NewParkingLotServiceMock()
		parkingLotService.On("GetParkedCarNumbersByColor", carColor).Return("KA-01-HH-1234, KA-01-HH-9999, KA-01-P-333", nil)
		parkingLotHandler := NewParkingLotHandler(parkingLotService)

		err := parkingLotHandler.GetParkedCarNumbersByColor(carColor)

		assert.Nil(t, err)
	})

	t.Run("get parked car numbers by empty car's color", func(t *testing.T) {
		parkingLotService := services.NewParkingLotServiceMock()
		parkingLotService.On("GetParkedCarNumbersByColor", "").Return("", errors.New("car's color is invalid"))
		parkingLotHandler := NewParkingLotHandler(parkingLotService)

		err := parkingLotHandler.GetParkedCarNumbersByColor("")

		assert.EqualError(t, err, "car's color is invalid")
	})

	t.Run("get parked car numbers by unavailable car's color", func(t *testing.T) {
		carColor := "White"
		parkingLotService := services.NewParkingLotServiceMock()
		parkingLotService.On("GetParkedCarNumbersByColor", carColor).Return("", errors.New("not found"))
		parkingLotHandler := NewParkingLotHandler(parkingLotService)

		err := parkingLotHandler.GetParkedCarNumbersByColor(carColor)

		assert.EqualError(t, err, "not found")
	})
}

func TestHandler_GetParkedSlotNumbersByColor(t *testing.T) {
	t.Run("get parked slot numbers by car's color", func(t *testing.T) {
		carColor := "White"
		parkingLotService := services.NewParkingLotServiceMock()
		parkingLotService.On("GetParkedSlotNumbersByColor", carColor).Return("1, 2, 4", nil)
		parkingLotHandler := NewParkingLotHandler(parkingLotService)

		err := parkingLotHandler.GetParkedSlotNumbersByColor(carColor)

		assert.Nil(t, err)
	})

	t.Run("get parked slot numbers by empty car's color", func(t *testing.T) {
		parkingLotService := services.NewParkingLotServiceMock()
		parkingLotService.On("GetParkedSlotNumbersByColor", "").Return("", errors.New("car's color is invalid"))
		parkingLotHandler := NewParkingLotHandler(parkingLotService)

		err := parkingLotHandler.GetParkedSlotNumbersByColor("")

		assert.EqualError(t, err, "car's color is invalid")
	})

	t.Run("get parked slot numbers by unavailable cars' color", func(t *testing.T) {
		carColor := "White"
		parkingLotService := services.NewParkingLotServiceMock()
		parkingLotService.On("GetParkedSlotNumbersByColor", carColor).Return("", errors.New("not found"))
		parkingLotHandler := NewParkingLotHandler(parkingLotService)

		err := parkingLotHandler.GetParkedSlotNumbersByColor(carColor)

		assert.EqualError(t, err, "not found")
	})
}

func TestHandler_GetParkedSlotNumberByCarNumber(t *testing.T) {
	t.Run("get parked slot numbers by car number", func(t *testing.T) {
		carNumber := "KA-01-HH-3141"
		parkingLotService := services.NewParkingLotServiceMock()
		parkingLotService.On("GetParkedSlotNumberByCarNumber", carNumber).Return(6, nil)
		parkingLotHandler := NewParkingLotHandler(parkingLotService)

		err := parkingLotHandler.GetParkedSlotNumberByCarNumber(carNumber)

		assert.Nil(t, err)
	})

	t.Run("get parked slot numbers by empty car number", func(t *testing.T) {
		parkingLotService := services.NewParkingLotServiceMock()
		parkingLotService.On("GetParkedSlotNumberByCarNumber", "").Return(0, errors.New("car's number is invalid"))
		parkingLotHandler := NewParkingLotHandler(parkingLotService)

		err := parkingLotHandler.GetParkedSlotNumberByCarNumber("")

		assert.EqualError(t, err, "car's number is invalid")
	})

	t.Run("get parked slot numbers by unavailable car number", func(t *testing.T) {
		carNumber := "KA-01-HH-3141"
		parkingLotService := services.NewParkingLotServiceMock()
		parkingLotService.On("GetParkedSlotNumberByCarNumber", carNumber).Return(0, errors.New("not found"))
		parkingLotHandler := NewParkingLotHandler(parkingLotService)

		err := parkingLotHandler.GetParkedSlotNumberByCarNumber(carNumber)

		assert.EqualError(t, err, "not found")
	})
}
