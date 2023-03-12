package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStore_CreateParkingLot(t *testing.T) {
	t.Run("create parking lot with 6 slots", func(t *testing.T) {
		parkingLotStore := NewParkingLotStore()
		parkingSlots, err := parkingLotStore.CreateParkingLot(6)
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
		parkingLotStore := NewParkingLotStore()
		parkingSlots, err := parkingLotStore.CreateParkingLot(0)
		assert.Nil(t, parkingSlots)
		assert.Error(t, err)
	})

	t.Run("create duplicated parking lot", func(t *testing.T) {
		parkingLotStore := NewParkingLotStore()
		parkingLotStore.parkingSlots = []*parkingSlot{
			{parkedCar: nil},
			{parkedCar: nil},
			{parkedCar: nil},
			{parkedCar: nil},
			{parkedCar: nil},
			{parkedCar: nil},
		}
		parkingSlots, err := parkingLotStore.CreateParkingLot(6)
		assert.Nil(t, parkingSlots)
		assert.Error(t, err)
	})
}

func TestStore_UpdateParkingLot(t *testing.T) {
	t.Run("update parking lot with car data", func(t *testing.T) {
		parkingLotStore := NewParkingLotStore()
		parkingLotStore.parkingSlots = []*parkingSlot{
			{parkedCar: nil},
			{parkedCar: nil},
			{parkedCar: nil},
			{parkedCar: nil},
			{parkedCar: nil},
			{parkedCar: nil},
		}
		parkingSlots, err := parkingLotStore.UpdateParkingLot(3, &car{
			registrationNumber: "KA-01-HH-2701",
			color:              "blue",
		})

		expected := []*parkingSlot{
			{parkedCar: nil},
			{parkedCar: nil},
			{parkedCar: nil},
			{parkedCar: &car{
				registrationNumber: "KA-01-HH-2701",
				color:              "blue",
			}},
			{parkedCar: nil},
			{parkedCar: nil},
		}

		assert.Equal(t, expected, parkingSlots)
		assert.Nil(t, err)
	})

	t.Run("update parking lot with nil data", func(t *testing.T) {
		parkingLotStore := NewParkingLotStore()
		parkingLotStore.parkingSlots = []*parkingSlot{
			{parkedCar: &car{
				registrationNumber: "KA-01-HH-1234",
				color:              "blue",
			}},
			{parkedCar: &car{
				registrationNumber: "KA-01-HH-9999",
				color:              "blue",
			}},
			{parkedCar: &car{
				registrationNumber: "KA-01-BB-0001",
				color:              "blue",
			}},
			{parkedCar: &car{
				registrationNumber: "KA-01-HH-7777",
				color:              "blue",
			}},
			{parkedCar: nil},
			{parkedCar: &car{
				registrationNumber: "KA-01-HH-2701",
				color:              "blue",
			}},
		}
		parkingSlots, err := parkingLotStore.UpdateParkingLot(3, nil)

		expected := []*parkingSlot{
			{parkedCar: &car{
				registrationNumber: "KA-01-HH-1234",
				color:              "blue",
			}},
			{parkedCar: &car{
				registrationNumber: "KA-01-HH-9999",
				color:              "blue",
			}},
			{parkedCar: &car{
				registrationNumber: "KA-01-BB-0001",
				color:              "blue",
			}},
			{parkedCar: nil},
			{parkedCar: nil},
			{parkedCar: &car{
				registrationNumber: "KA-01-HH-2701",
				color:              "blue",
			}},
		}

		assert.Equal(t, expected, parkingSlots)
		assert.Nil(t, err)
	})

	t.Run("update parking lot with negative index", func(t *testing.T) {
		parkingLotStore := NewParkingLotStore()
		parkingLotStore.parkingSlots = []*parkingSlot{
			{parkedCar: nil},
			{parkedCar: nil},
			{parkedCar: nil},
			{parkedCar: nil},
			{parkedCar: nil},
			{parkedCar: nil},
		}
		parkingSlots, err := parkingLotStore.UpdateParkingLot(-1, nil)
		assert.Nil(t, parkingSlots)
		assert.Error(t, err)
	})

	t.Run("update invalid parking lot", func(t *testing.T) {
		parkingLotStore := NewParkingLotStore()
		parkingLotStore.parkingSlots = []*parkingSlot{
			{parkedCar: nil},
			{parkedCar: nil},
			{parkedCar: nil},
			{parkedCar: nil},
			{parkedCar: nil},
			{parkedCar: nil},
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
		parkingLotStore.parkingSlots = []*parkingSlot{
			{parkedCar: nil},
			{parkedCar: nil},
			{parkedCar: nil},
			{parkedCar: &car{
				registrationNumber: "KA-01-HH-2701",
				color:              "blue",
			}},
			{parkedCar: nil},
			{parkedCar: nil},
		}
		parkingSlotDetail, err := parkingLotStore.GetParkingSlotDetail(3)
		expected := &parkingSlot{parkedCar: &car{
			registrationNumber: "KA-01-HH-2701",
			color:              "blue",
		}}
		assert.Equal(t, expected, parkingSlotDetail)
		assert.Nil(t, err)
	})

	t.Run("get parking lot detail by negative index", func(t *testing.T) {
		parkingLotStore := NewParkingLotStore()
		parkingLotStore.parkingSlots = []*parkingSlot{
			{parkedCar: nil},
			{parkedCar: nil},
			{parkedCar: nil},
			{parkedCar: &car{
				registrationNumber: "KA-01-HH-2701",
				color:              "blue",
			}},
			{parkedCar: nil},
			{parkedCar: nil},
		}
		parkingSlotDetail, err := parkingLotStore.GetParkingSlotDetail(-1)

		assert.Nil(t, parkingSlotDetail)
		assert.Error(t, err)
	})

	t.Run("get unavailable parking lot detail", func(t *testing.T) {
		parkingLotStore := NewParkingLotStore()
		parkingLotStore.parkingSlots = []*parkingSlot{
			{parkedCar: nil},
			{parkedCar: nil},
			{parkedCar: nil},
			{parkedCar: &car{
				registrationNumber: "KA-01-HH-2701",
				color:              "blue",
			}},
			{parkedCar: nil},
			{parkedCar: nil},
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
