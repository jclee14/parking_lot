package stores

import (
	"parking_lot/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStore_CreateParkingLot(t *testing.T) {
	t.Run("create parking lot with 6 slots", func(t *testing.T) {
		parkingLotStore := NewParkingLotStore()
		parkingSlots, err := parkingLotStore.CreateParkingLot(6)
		expected := []*models.ParkingSlot{
			{ParkedCar: nil},
			{ParkedCar: nil},
			{ParkedCar: nil},
			{ParkedCar: nil},
			{ParkedCar: nil},
			{ParkedCar: nil},
		}
		assert.Equal(t, expected, parkingSlots)
		assert.Nil(t, err)
	})

	t.Run("create parking lot with 0 slot", func(t *testing.T) {
		parkingLotStore := NewParkingLotStore()
		parkingSlots, err := parkingLotStore.CreateParkingLot(0)
		assert.Nil(t, parkingSlots)
		assert.Error(t, err)
	})

	t.Run("create duplicated parking lot", func(t *testing.T) {
		parkingLotStore := NewParkingLotStore()
		parkingLotStore.parkingSlots = []*models.ParkingSlot{
			{ParkedCar: nil},
			{ParkedCar: nil},
			{ParkedCar: nil},
			{ParkedCar: nil},
			{ParkedCar: nil},
			{ParkedCar: nil},
		}
		parkingSlots, err := parkingLotStore.CreateParkingLot(6)
		assert.Nil(t, parkingSlots)
		assert.Error(t, err)
	})
}

func TestStore_UpdateParkingLot(t *testing.T) {
	t.Run("update parking lot with car data", func(t *testing.T) {
		parkingLotStore := NewParkingLotStore()
		parkingLotStore.parkingSlots = []*models.ParkingSlot{
			{ParkedCar: nil},
			{ParkedCar: nil},
			{ParkedCar: nil},
			{ParkedCar: nil},
			{ParkedCar: nil},
			{ParkedCar: nil},
		}
		parkingSlots, err := parkingLotStore.UpdateParkingLot(3, &models.Car{
			RegistrationNumber: "KA-01-HH-2701",
			Color:              "blue",
		})

		expected := []*models.ParkingSlot{
			{ParkedCar: nil},
			{ParkedCar: nil},
			{ParkedCar: nil},
			{ParkedCar: &models.Car{
				RegistrationNumber: "KA-01-HH-2701",
				Color:              "blue",
			}},
			{ParkedCar: nil},
			{ParkedCar: nil},
		}

		assert.Equal(t, expected, parkingSlots)
		assert.Nil(t, err)
	})

	t.Run("update parking lot with nil data", func(t *testing.T) {
		parkingLotStore := NewParkingLotStore()
		parkingLotStore.parkingSlots = []*models.ParkingSlot{
			{ParkedCar: &models.Car{
				RegistrationNumber: "KA-01-HH-1234",
				Color:              "blue",
			}},
			{ParkedCar: &models.Car{
				RegistrationNumber: "KA-01-HH-9999",
				Color:              "blue",
			}},
			{ParkedCar: &models.Car{
				RegistrationNumber: "KA-01-BB-0001",
				Color:              "blue",
			}},
			{ParkedCar: &models.Car{
				RegistrationNumber: "KA-01-HH-7777",
				Color:              "blue",
			}},
			{ParkedCar: nil},
			{ParkedCar: &models.Car{
				RegistrationNumber: "KA-01-HH-2701",
				Color:              "blue",
			}},
		}
		parkingSlots, err := parkingLotStore.UpdateParkingLot(3, nil)

		expected := []*models.ParkingSlot{
			{ParkedCar: &models.Car{
				RegistrationNumber: "KA-01-HH-1234",
				Color:              "blue",
			}},
			{ParkedCar: &models.Car{
				RegistrationNumber: "KA-01-HH-9999",
				Color:              "blue",
			}},
			{ParkedCar: &models.Car{
				RegistrationNumber: "KA-01-BB-0001",
				Color:              "blue",
			}},
			{ParkedCar: nil},
			{ParkedCar: nil},
			{ParkedCar: &models.Car{
				RegistrationNumber: "KA-01-HH-2701",
				Color:              "blue",
			}},
		}

		assert.Equal(t, expected, parkingSlots)
		assert.Nil(t, err)
	})

	t.Run("update parking lot with negative index", func(t *testing.T) {
		parkingLotStore := NewParkingLotStore()
		parkingLotStore.parkingSlots = []*models.ParkingSlot{
			{ParkedCar: nil},
			{ParkedCar: nil},
			{ParkedCar: nil},
			{ParkedCar: nil},
			{ParkedCar: nil},
			{ParkedCar: nil},
		}
		parkingSlots, err := parkingLotStore.UpdateParkingLot(-1, nil)
		assert.Nil(t, parkingSlots)
		assert.Error(t, err)
	})

	t.Run("update invalid parking lot", func(t *testing.T) {
		parkingLotStore := NewParkingLotStore()
		parkingLotStore.parkingSlots = []*models.ParkingSlot{
			{ParkedCar: nil},
			{ParkedCar: nil},
			{ParkedCar: nil},
			{ParkedCar: nil},
			{ParkedCar: nil},
			{ParkedCar: nil},
		}
		parkingSlots, err := parkingLotStore.UpdateParkingLot(6, nil)
		assert.Nil(t, parkingSlots)
		assert.Error(t, err)
	})

	t.Run("update parking lot before it created", func(t *testing.T) {
		parkingLotStore := NewParkingLotStore()
		parkingSlots, err := parkingLotStore.UpdateParkingLot(6, nil)
		assert.Nil(t, parkingSlots)
		assert.Error(t, err)
	})
}

func TestStore_GetParkingSlotDetail(t *testing.T) {
	t.Run("get parking lot detail", func(t *testing.T) {
		parkingLotStore := NewParkingLotStore()
		parkingLotStore.parkingSlots = []*models.ParkingSlot{
			{ParkedCar: nil},
			{ParkedCar: nil},
			{ParkedCar: nil},
			{ParkedCar: &models.Car{
				RegistrationNumber: "KA-01-HH-2701",
				Color:              "blue",
			}},
			{ParkedCar: nil},
			{ParkedCar: nil},
		}
		parkingSlotDetail, err := parkingLotStore.GetParkingSlotDetail(3)
		expected := &models.ParkingSlot{ParkedCar: &models.Car{
			RegistrationNumber: "KA-01-HH-2701",
			Color:              "blue",
		}}
		assert.Equal(t, expected, parkingSlotDetail)
		assert.Nil(t, err)
	})

	t.Run("get parking lot detail by negative index", func(t *testing.T) {
		parkingLotStore := NewParkingLotStore()
		parkingLotStore.parkingSlots = []*models.ParkingSlot{
			{ParkedCar: nil},
			{ParkedCar: nil},
			{ParkedCar: nil},
			{ParkedCar: &models.Car{
				RegistrationNumber: "KA-01-HH-2701",
				Color:              "blue",
			}},
			{ParkedCar: nil},
			{ParkedCar: nil},
		}
		parkingSlotDetail, err := parkingLotStore.GetParkingSlotDetail(-1)

		assert.Nil(t, parkingSlotDetail)
		assert.Error(t, err)
	})

	t.Run("get unavailable parking lot detail", func(t *testing.T) {
		parkingLotStore := NewParkingLotStore()
		parkingLotStore.parkingSlots = []*models.ParkingSlot{
			{ParkedCar: nil},
			{ParkedCar: nil},
			{ParkedCar: nil},
			{ParkedCar: &models.Car{
				RegistrationNumber: "KA-01-HH-2701",
				Color:              "blue",
			}},
			{ParkedCar: nil},
			{ParkedCar: nil},
		}
		parkingSlotDetail, err := parkingLotStore.GetParkingSlotDetail(6)

		assert.Nil(t, parkingSlotDetail)
		assert.Error(t, err)
	})

	t.Run("get parking lot detail before it created", func(t *testing.T) {
		parkingLotStore := NewParkingLotStore()
		parkingSlotDetail, err := parkingLotStore.GetParkingSlotDetail(3)

		assert.Nil(t, parkingSlotDetail)
		assert.Error(t, err)
	})
}
