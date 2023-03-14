package services

import (
	"errors"
	"parking_lot/models"
	"parking_lot/stores"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestService_CreateParkingLot(t *testing.T) {
	t.Run("create parking lot with 6 slots", func(t *testing.T) {
		parkingLotStore := stores.NewParkingLotStoreMock()
		parkingLotStore.On("CreateParkingLot", 6).Return(
			[]*models.ParkingSlot{
				{ParkedCar: nil},
				{ParkedCar: nil},
				{ParkedCar: nil},
				{ParkedCar: nil},
				{ParkedCar: nil},
				{ParkedCar: nil},
			}, nil)
		parkingLotService := NewParkingLotService(parkingLotStore)

		parkingSlots, err := parkingLotService.CreateParkingLot(6)
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
		parkingLotStore := stores.NewParkingLotStoreMock()
		parkingLotStore.On("CreateParkingLot", 0).Return(nil, errors.New("slot amount must be greater than 0"))
		parkingLotService := NewParkingLotService(parkingLotStore)

		parkingSlots, err := parkingLotService.CreateParkingLot(0)

		assert.Nil(t, parkingSlots)
		assert.Error(t, err)
	})
}

func TestService_Park(t *testing.T) {
	type testCase struct {
		testName               string
		carNumber              string
		carColor               string
		parkingSlotData        []*models.ParkingSlot
		updatedParkingSlotData []*models.ParkingSlot
		toParkSlotNumber       int
		parkedSlotNumber       int
	}

	testCases := []testCase{
		{
			testName:  "register first car to parking lot",
			carNumber: "KA-01-HH-2701", carColor: "blue", toParkSlotNumber: 0, parkedSlotNumber: 1,
			parkingSlotData:        []*models.ParkingSlot{{ParkedCar: nil}, {ParkedCar: nil}, {ParkedCar: nil}, {ParkedCar: nil}, {ParkedCar: nil}, {ParkedCar: nil}},
			updatedParkingSlotData: []*models.ParkingSlot{{ParkedCar: &models.Car{RegistrationNumber: "KA-01-HH-2701", Color: "blue"}}, {ParkedCar: nil}, {ParkedCar: nil}, {ParkedCar: nil}, {ParkedCar: nil}, {ParkedCar: nil}},
		},
		{
			testName:  "register car to parking lot number 3",
			carNumber: "park KA-01-HH-7777", carColor: "red", toParkSlotNumber: 2, parkedSlotNumber: 3,
			parkingSlotData:        []*models.ParkingSlot{{ParkedCar: &models.Car{RegistrationNumber: "KA-01-HH-2701", Color: "blue"}}, {ParkedCar: &models.Car{RegistrationNumber: "KA-01-HH-2702", Color: "blue"}}, {ParkedCar: nil}, {ParkedCar: nil}, {ParkedCar: &models.Car{RegistrationNumber: "KA-01-HH-2705", Color: "blue"}}, {ParkedCar: nil}},
			updatedParkingSlotData: []*models.ParkingSlot{{ParkedCar: &models.Car{RegistrationNumber: "KA-01-HH-2701", Color: "blue"}}, {ParkedCar: &models.Car{RegistrationNumber: "KA-01-HH-2702", Color: "blue"}}, {ParkedCar: &models.Car{RegistrationNumber: "park KA-01-HH-7777", Color: "red"}}, {ParkedCar: nil}, {ParkedCar: &models.Car{RegistrationNumber: "KA-01-HH-2705", Color: "blue"}}, {ParkedCar: nil}},
		},
		{
			testName:  "register last car to parking lot",
			carNumber: "KA-01-HH-9999", carColor: "white", toParkSlotNumber: 5, parkedSlotNumber: 6,
			parkingSlotData:        []*models.ParkingSlot{{ParkedCar: &models.Car{RegistrationNumber: "KA-01-HH-2701", Color: "blue"}}, {ParkedCar: &models.Car{RegistrationNumber: "KA-01-HH-2702", Color: "blue"}}, {ParkedCar: &models.Car{RegistrationNumber: "KA-01-HH-2703", Color: "blue"}}, {ParkedCar: &models.Car{RegistrationNumber: "KA-01-HH-2704", Color: "blue"}}, {ParkedCar: &models.Car{RegistrationNumber: "KA-01-HH-2705", Color: "blue"}}, {ParkedCar: nil}},
			updatedParkingSlotData: []*models.ParkingSlot{{ParkedCar: &models.Car{RegistrationNumber: "KA-01-HH-2701", Color: "blue"}}, {ParkedCar: &models.Car{RegistrationNumber: "KA-01-HH-2702", Color: "blue"}}, {ParkedCar: &models.Car{RegistrationNumber: "KA-01-HH-2703", Color: "blue"}}, {ParkedCar: &models.Car{RegistrationNumber: "KA-01-HH-2704", Color: "blue"}}, {ParkedCar: &models.Car{RegistrationNumber: "KA-01-HH-2705", Color: "blue"}}, {ParkedCar: &models.Car{RegistrationNumber: "KA-01-HH-9999", Color: "white"}}},
		},
	}

	for _, caseData := range testCases {
		t.Run(caseData.testName, func(t *testing.T) {
			parkingLotStore := stores.NewParkingLotStoreMock()
			parkingLotStore.On("GetParkingSlots").Return(caseData.parkingSlotData, nil)
			parkingLotStore.On("UpdateParkingLot", caseData.toParkSlotNumber, &models.Car{RegistrationNumber: caseData.carNumber, Color: caseData.carColor}).Return(caseData.updatedParkingSlotData, nil)
			parkingLotService := NewParkingLotService(parkingLotStore)

			parkedSlotNumber, err := parkingLotService.Park(caseData.carNumber, caseData.carColor)
			expected := caseData.parkedSlotNumber

			assert.Equal(t, expected, parkedSlotNumber)
			assert.Nil(t, err)
		})
	}

	t.Run("register car with empty car's color", func(t *testing.T) {
		parkingLotStore := stores.NewParkingLotStoreMock()
		parkingLotStore.On("GetParkingSlots").Return([]*models.ParkingSlot{{ParkedCar: nil}, {ParkedCar: nil}, {ParkedCar: nil}, {ParkedCar: nil}, {ParkedCar: nil}, {ParkedCar: nil}}, nil)
		parkingLotService := NewParkingLotService(parkingLotStore)

		parkedSlotNumber, err := parkingLotService.Park("KA-01-HH-2701", "")
		expected := 0

		assert.Equal(t, expected, parkedSlotNumber)
		assert.Error(t, err)
	})
	t.Run("register car with empty car's number", func(t *testing.T) {
		parkingLotStore := stores.NewParkingLotStoreMock()
		parkingLotStore.On("GetParkingSlots").Return([]*models.ParkingSlot{{ParkedCar: nil}, {ParkedCar: nil}, {ParkedCar: nil}, {ParkedCar: nil}, {ParkedCar: nil}, {ParkedCar: nil}}, nil)
		parkingLotService := NewParkingLotService(parkingLotStore)

		parkedSlotNumber, err := parkingLotService.Park("", "black")
		expected := 0

		assert.Equal(t, expected, parkedSlotNumber)
		assert.Error(t, err)
	})

	t.Run("register car before create parking lot", func(t *testing.T) {
		parkingLotStore := stores.NewParkingLotStoreMock()
		parkingLotStore.On("GetParkingSlots").Return(nil, errors.New("parking lot is not created yet"))
		parkingLotService := NewParkingLotService(parkingLotStore)

		parkedSlotNumber, err := parkingLotService.Park("KA-01-HH-2701", "black")
		expected := 0

		assert.Equal(t, expected, parkedSlotNumber)
		assert.Error(t, err)
	})

	t.Run("register car to full parking lot", func(t *testing.T) {
		carNumber := "KA-01-HH-2701"
		carColor := "black"

		parkingLotStore := stores.NewParkingLotStoreMock()
		parkingLotStore.On("GetParkingSlots").Return([]*models.ParkingSlot{{ParkedCar: &models.Car{RegistrationNumber: "KA-01-HH-2701", Color: "blue"}}, {ParkedCar: &models.Car{RegistrationNumber: "KA-01-HH-2702", Color: "blue"}}, {ParkedCar: &models.Car{RegistrationNumber: "KA-01-HH-2703", Color: "blue"}}, {ParkedCar: &models.Car{RegistrationNumber: "KA-01-HH-2704", Color: "blue"}}, {ParkedCar: &models.Car{RegistrationNumber: "KA-01-HH-2705", Color: "blue"}}, {ParkedCar: &models.Car{RegistrationNumber: "KA-01-HH-9999", Color: "white"}}}, nil)
		parkingLotService := NewParkingLotService(parkingLotStore)

		parkedSlotNumber, err := parkingLotService.Park(carNumber, carColor)
		expected := 0

		assert.Equal(t, expected, parkedSlotNumber)
		assert.Error(t, err)
	})

	t.Run("register car and get error when update", func(t *testing.T) {
		carNumber := "KA-01-HH-2701"
		carColor := "black"

		parkingLotStore := stores.NewParkingLotStoreMock()
		parkingLotStore.On("GetParkingSlots").Return([]*models.ParkingSlot{{ParkedCar: &models.Car{RegistrationNumber: "KA-01-HH-2701", Color: "blue"}}, {ParkedCar: &models.Car{RegistrationNumber: "KA-01-HH-2702", Color: "blue"}}, {ParkedCar: &models.Car{RegistrationNumber: "KA-01-HH-2703", Color: "blue"}}, {ParkedCar: &models.Car{RegistrationNumber: "KA-01-HH-2704", Color: "blue"}}, {ParkedCar: &models.Car{RegistrationNumber: "KA-01-HH-2705", Color: "blue"}}, {ParkedCar: nil}}, nil)
		parkingLotStore.On("UpdateParkingLot", 5, &models.Car{RegistrationNumber: carNumber, Color: carColor}).Return(nil, errors.New("internal server error"))
		parkingLotService := NewParkingLotService(parkingLotStore)

		parkedSlotNumber, err := parkingLotService.Park(carNumber, carColor)
		expected := 0

		assert.Equal(t, expected, parkedSlotNumber)
		assert.Error(t, err)
	})
}

func TestService_CarLeaving(t *testing.T) {
	type testCase struct {
		testName               string
		slotNumber             int
		toUpdateData           *models.Car
		parkingSlotDetail      *models.ParkingSlot
		parkingSlotsData       []*models.ParkingSlot
		updatedParkingSlotData []*models.ParkingSlot
	}

	testCases := []testCase{
		{
			testName:               "leave slot number 1",
			slotNumber:             1,
			toUpdateData:           nil,
			parkingSlotDetail:      &models.ParkingSlot{ParkedCar: &models.Car{RegistrationNumber: "KA-01-HH-2701", Color: "blue"}},
			parkingSlotsData:       []*models.ParkingSlot{{ParkedCar: &models.Car{RegistrationNumber: "KA-01-HH-2701", Color: "blue"}}, {ParkedCar: nil}, {ParkedCar: nil}, {ParkedCar: nil}, {ParkedCar: nil}, {ParkedCar: nil}},
			updatedParkingSlotData: []*models.ParkingSlot{{ParkedCar: nil}, {ParkedCar: nil}, {ParkedCar: nil}, {ParkedCar: nil}, {ParkedCar: nil}, {ParkedCar: nil}},
		},
		{
			testName:               "leave slot number 6",
			slotNumber:             6,
			toUpdateData:           nil,
			parkingSlotDetail:      &models.ParkingSlot{ParkedCar: &models.Car{RegistrationNumber: "KA-01-HH-9999", Color: "white"}},
			parkingSlotsData:       []*models.ParkingSlot{{ParkedCar: &models.Car{RegistrationNumber: "KA-01-HH-2701", Color: "blue"}}, {ParkedCar: &models.Car{RegistrationNumber: "KA-01-HH-2702", Color: "blue"}}, {ParkedCar: &models.Car{RegistrationNumber: "KA-01-HH-2703", Color: "blue"}}, {ParkedCar: &models.Car{RegistrationNumber: "KA-01-HH-2704", Color: "blue"}}, {ParkedCar: &models.Car{RegistrationNumber: "KA-01-HH-2705", Color: "blue"}}, {ParkedCar: &models.Car{RegistrationNumber: "KA-01-HH-9999", Color: "white"}}},
			updatedParkingSlotData: []*models.ParkingSlot{{ParkedCar: &models.Car{RegistrationNumber: "KA-01-HH-2701", Color: "blue"}}, {ParkedCar: &models.Car{RegistrationNumber: "KA-01-HH-2702", Color: "blue"}}, {ParkedCar: &models.Car{RegistrationNumber: "KA-01-HH-2703", Color: "blue"}}, {ParkedCar: &models.Car{RegistrationNumber: "KA-01-HH-2704", Color: "blue"}}, {ParkedCar: &models.Car{RegistrationNumber: "KA-01-HH-2705", Color: "blue"}}, {ParkedCar: nil}},
		},
		{
			testName:               "leave slot number 3",
			slotNumber:             3,
			toUpdateData:           nil,
			parkingSlotDetail:      &models.ParkingSlot{ParkedCar: &models.Car{RegistrationNumber: "KA-01-HH-2703", Color: "blue"}},
			parkingSlotsData:       []*models.ParkingSlot{{ParkedCar: &models.Car{RegistrationNumber: "KA-01-HH-2701", Color: "blue"}}, {ParkedCar: &models.Car{RegistrationNumber: "KA-01-HH-2702", Color: "blue"}}, {ParkedCar: &models.Car{RegistrationNumber: "KA-01-HH-2703", Color: "blue"}}, {ParkedCar: nil}, {ParkedCar: &models.Car{RegistrationNumber: "KA-01-HH-2705", Color: "blue"}}, {ParkedCar: nil}},
			updatedParkingSlotData: []*models.ParkingSlot{{ParkedCar: &models.Car{RegistrationNumber: "KA-01-HH-2701", Color: "blue"}}, {ParkedCar: &models.Car{RegistrationNumber: "KA-01-HH-2702", Color: "blue"}}, {ParkedCar: nil}, {ParkedCar: nil}, {ParkedCar: &models.Car{RegistrationNumber: "KA-01-HH-2705", Color: "blue"}}, {ParkedCar: nil}},
		},
	}

	for _, caseData := range testCases {
		t.Run(caseData.testName, func(t *testing.T) {
			parkingLotStore := stores.NewParkingLotStoreMock()
			parkingLotStore.On("GetParkingSlotDetail", caseData.slotNumber-1).Return(caseData.parkingSlotDetail, nil)
			parkingLotStore.On("UpdateParkingLot", caseData.slotNumber-1, caseData.toUpdateData).Return(caseData.updatedParkingSlotData, nil)
			parkingLotService := NewParkingLotService(parkingLotStore)

			err := parkingLotService.Leave(caseData.slotNumber)
			assert.Nil(t, err)
		})
	}

	t.Run("leave nil parking lot", func(t *testing.T) {
		leaveSlot := 1
		parkingLotStore := stores.NewParkingLotStoreMock()
		parkingLotStore.On("GetParkingSlotDetail", leaveSlot-1).Return(nil, nil)
		parkingLotService := NewParkingLotService(parkingLotStore)

		err := parkingLotService.Leave(leaveSlot)
		assert.Error(t, err)
	})

	t.Run("leave unavailable parking lot", func(t *testing.T) {
		leaveSlot := 1
		parkingLotStore := stores.NewParkingLotStoreMock()
		parkingLotStore.On("GetParkingSlotDetail", leaveSlot-1).Return(nil, errors.New("parking lot is not created yet"))
		parkingLotService := NewParkingLotService(parkingLotStore)

		err := parkingLotService.Leave(leaveSlot)
		assert.Error(t, err)
	})

	t.Run("leave unoccupied parking lot", func(t *testing.T) {
		leaveSlot := 1
		parkingLotStore := stores.NewParkingLotStoreMock()
		parkingLotStore.On("GetParkingSlotDetail", leaveSlot-1).Return(&models.ParkingSlot{ParkedCar: nil}, nil)
		parkingLotService := NewParkingLotService(parkingLotStore)

		err := parkingLotService.Leave(leaveSlot)
		assert.Error(t, err)
	})

	t.Run("leave car and get error when update", func(t *testing.T) {
		leaveSlot := 1
		var toUpdateData *models.Car = nil

		parkingLotStore := stores.NewParkingLotStoreMock()
		parkingLotStore.On("GetParkingSlotDetail", leaveSlot-1).Return(&models.ParkingSlot{ParkedCar: &models.Car{RegistrationNumber: "KA-01-HH-2703", Color: "blue"}}, nil)
		parkingLotStore.On("UpdateParkingLot", leaveSlot-1, toUpdateData).Return(nil, errors.New("internal server error"))
		parkingLotService := NewParkingLotService(parkingLotStore)

		err := parkingLotService.Leave(leaveSlot)
		assert.Error(t, err)
	})
}

func TestService_GetStatus(t *testing.T) {
	type testCase struct {
		testName         string
		parkingSlotsData []*models.ParkingSlot
		expected         string
	}

	testCases := []testCase{
		{
			testName:         "get full parking lot status",
			parkingSlotsData: []*models.ParkingSlot{{ParkedCar: &models.Car{RegistrationNumber: "KA-01-HH-2701", Color: "blue"}}, {ParkedCar: &models.Car{RegistrationNumber: "KA-01-HH-2702", Color: "blue"}}, {ParkedCar: &models.Car{RegistrationNumber: "KA-01-HH-2703", Color: "red"}}, {ParkedCar: &models.Car{RegistrationNumber: "KA-01-HH-2704", Color: "yellow"}}, {ParkedCar: &models.Car{RegistrationNumber: "KA-01-HH-2705", Color: "blue"}}, {ParkedCar: &models.Car{RegistrationNumber: "KA-01-HH-2706", Color: "white"}}},
			expected:         "1\t\tKA-01-HH-2701\t\tBlue\n2\t\tKA-01-HH-2702\t\tBlue\n3\t\tKA-01-HH-2703\t\tRed\n4\t\tKA-01-HH-2704\t\tYellow\n5\t\tKA-01-HH-2705\t\tBlue\n6\t\tKA-01-HH-2706\t\tWhite\n",
		},
		{
			testName:         "get parking lot status",
			parkingSlotsData: []*models.ParkingSlot{{ParkedCar: nil}, {ParkedCar: &models.Car{RegistrationNumber: "KA-01-HH-2702", Color: "blue"}}, {ParkedCar: nil}, {ParkedCar: nil}, {ParkedCar: &models.Car{RegistrationNumber: "KA-01-HH-2705", Color: "blue"}}, {ParkedCar: &models.Car{RegistrationNumber: "KA-01-HH-2706", Color: "white"}}},
			expected:         "2\t\tKA-01-HH-2702\t\tBlue\n5\t\tKA-01-HH-2705\t\tBlue\n6\t\tKA-01-HH-2706\t\tWhite\n",
		},
		{
			testName:         "get empty parking lot status",
			parkingSlotsData: []*models.ParkingSlot{{ParkedCar: nil}, {ParkedCar: nil}, {ParkedCar: nil}, {ParkedCar: nil}, {ParkedCar: nil}, {ParkedCar: nil}},
			expected:         "",
		},
		{
			testName:         "get nil parking lot status",
			parkingSlotsData: nil,
			expected:         "",
		},
	}

	for _, testData := range testCases {
		t.Run(testData.testName, func(t *testing.T) {
			parkingLotStore := stores.NewParkingLotStoreMock()
			parkingLotStore.On("GetParkingSlots").Return(testData.parkingSlotsData, nil)
			parkingLotService := NewParkingLotService(parkingLotStore)
			expected := testData.expected

			status, err := parkingLotService.GetStatus()
			assert.Equal(t, expected, status)
			assert.Nil(t, err)
		})
	}

	t.Run("get unavailable parking lot status", func(t *testing.T) {
		parkingLotStore := stores.NewParkingLotStoreMock()
		parkingLotStore.On("GetParkingSlots").Return(nil, errors.New("parking lot is not created yet"))
		parkingLotService := NewParkingLotService(parkingLotStore)

		status, err := parkingLotService.GetStatus()
		assert.Equal(t, "", status)
		assert.Error(t, err)
	})
}

func TestService_GetParkedCarNumbersByColor(t *testing.T) {
	t.Run("get parked car numbers by empty color string", func(t *testing.T) {
		parkingLotStore := stores.NewParkingLotStoreMock()
		parkingLotService := NewParkingLotService(parkingLotStore)
		data, err := parkingLotService.GetParkedCarNumbersByColor("")
		assert.Equal(t, "", data)
		assert.Error(t, err)
	})

	t.Run("get empty parked car numbers by color", func(t *testing.T) {
		parkingLotStore := stores.NewParkingLotStoreMock()
		parkingLotService := NewParkingLotService(parkingLotStore)
		parkingLotService.carNumbersByColor = map[string]map[string]struct{}{
			"blue": {
				"KA-01-HH-2701": struct{}{},
				"KA-01-HH-2702": struct{}{},
			},
			"red": {
				"KA-01-HH-2703": struct{}{},
				"":              struct{}{},
				"KA-01-HH-2704": struct{}{},
			},
		}
		data, err := parkingLotService.GetParkedCarNumbersByColor("white")
		assert.Equal(t, "", data)
		assert.Error(t, err)
	})

	t.Run("get parked car numbers by color", func(t *testing.T) {
		parkingLotStore := stores.NewParkingLotStoreMock()
		parkingLotService := NewParkingLotService(parkingLotStore)
		parkingLotService.carNumbersByColor = map[string]map[string]struct{}{
			"blue": {
				"KA-01-HH-2701": struct{}{},
				"KA-01-HH-2702": struct{}{},
			},
			"red": {
				"KA-01-HH-2703": struct{}{},
				"":              struct{}{},
			},
		}
		data, err := parkingLotService.GetParkedCarNumbersByColor("red")
		expected := "KA-01-HH-2703"
		assert.Equal(t, expected, data)
		assert.Nil(t, err)
	})
}

func TestService_GetParkedSlotNumbersByColor(t *testing.T) {
	t.Run("get parked slot numbers by empty color string", func(t *testing.T) {
		parkingLotStore := stores.NewParkingLotStoreMock()
		parkingLotService := NewParkingLotService(parkingLotStore)
		data, err := parkingLotService.GetParkedSlotNumbersByColor("")
		assert.Equal(t, "", data)
		assert.Error(t, err)
	})

	t.Run("get empty parked slot numbers by color", func(t *testing.T) {
		parkingLotStore := stores.NewParkingLotStoreMock()
		parkingLotService := NewParkingLotService(parkingLotStore)
		parkingLotService.slotNumbersByColor = map[string]map[int]struct{}{
			"blue": {
				5: struct{}{},
				3: struct{}{},
			},
			"red": {
				0: struct{}{},
				4: struct{}{},
			},
		}
		data, err := parkingLotService.GetParkedSlotNumbersByColor("white")
		assert.Equal(t, "", data)
		assert.Error(t, err)
	})

	t.Run("get parked slot numbers by color", func(t *testing.T) {
		parkingLotStore := stores.NewParkingLotStoreMock()
		parkingLotService := NewParkingLotService(parkingLotStore)
		parkingLotService.slotNumbersByColor = map[string]map[int]struct{}{
			"blue": {
				5: struct{}{},
				3: struct{}{},
			},
			"red": {
				0: struct{}{},
			},
		}
		data, err := parkingLotService.GetParkedSlotNumbersByColor("red")
		expected := "1"
		assert.Equal(t, expected, data)
		assert.Nil(t, err)
	})
}

func TestService_GetParkedSlotNumberByCarNumber(t *testing.T) {
	t.Run("get parked slot number by empty car number", func(t *testing.T) {
		parkingLotStore := stores.NewParkingLotStoreMock()
		parkingLotService := NewParkingLotService(parkingLotStore)
		data, err := parkingLotService.GetParkedSlotNumberByCarNumber("")
		assert.Equal(t, 0, data)
		assert.Error(t, err)
	})

	t.Run("get parked slot number by unavailable car number", func(t *testing.T) {
		parkingLotStore := stores.NewParkingLotStoreMock()
		parkingLotService := NewParkingLotService(parkingLotStore)
		parkingLotService.slotNumberByCarNumber = map[string]int{
			"KA-01-HH-2703": 0,
			"KA-01-HH-2705": 3,
			"KA-01-HH-2708": 4,
		}
		data, err := parkingLotService.GetParkedSlotNumberByCarNumber("KA-01-HH-2709")
		assert.Equal(t, 0, data)
		assert.Error(t, err)
	})

	t.Run("get parked slot number by car number", func(t *testing.T) {
		parkingLotStore := stores.NewParkingLotStoreMock()
		parkingLotService := NewParkingLotService(parkingLotStore)
		parkingLotService.slotNumberByCarNumber = map[string]int{
			"KA-01-HH-2703": 0,
			"KA-01-HH-2705": 3,
			"KA-01-HH-2708": 4,
		}
		data, err := parkingLotService.GetParkedSlotNumberByCarNumber("KA-01-HH-2705")
		expected := 4
		assert.Equal(t, expected, data)
		assert.Nil(t, err)
	})
}

func TestService_addCarNumberToColorCache(t *testing.T) {
	t.Run("add empty car number to color cache", func(t *testing.T) {
		parkingLotStore := stores.NewParkingLotStoreMock()
		parkingLotService := NewParkingLotService(parkingLotStore)
		err := parkingLotService.addCarNumberToColorCache("", "white")
		assert.Error(t, err)
	})
	t.Run("add car number to empty color cache", func(t *testing.T) {
		parkingLotStore := stores.NewParkingLotStoreMock()
		parkingLotService := NewParkingLotService(parkingLotStore)
		err := parkingLotService.addCarNumberToColorCache("KA-01-HH-2708", "")
		assert.Error(t, err)
	})

	t.Run("add car number to new color cache", func(t *testing.T) {
		parkingLotStore := stores.NewParkingLotStoreMock()
		parkingLotService := NewParkingLotService(parkingLotStore)
		err := parkingLotService.addCarNumberToColorCache("KA-01-HH-2708", "white")
		expected := map[string]map[string]struct{}{
			"white": {
				"KA-01-HH-2708": struct{}{},
			},
		}
		assert.Equal(t, expected, parkingLotService.carNumbersByColor)
		assert.Nil(t, err)
	})

	t.Run("add car number to exist color cache", func(t *testing.T) {
		parkingLotStore := stores.NewParkingLotStoreMock()
		parkingLotService := NewParkingLotService(parkingLotStore)
		parkingLotService.carNumbersByColor = map[string]map[string]struct{}{
			"white": {
				"KA-01-HH-2708": struct{}{},
			},
		}
		err := parkingLotService.addCarNumberToColorCache("KA-01-HH-2709", "white")
		expected := map[string]map[string]struct{}{
			"white": {
				"KA-01-HH-2708": struct{}{},
				"KA-01-HH-2709": struct{}{},
			},
		}
		if !reflect.DeepEqual(expected, parkingLotService.carNumbersByColor) {
			t.Errorf("mismatch, got: %v wanted: %v", expected, parkingLotService.carNumbersByColor)
		}
		assert.Nil(t, err)
	})
}

func TestService_removeCarNumberFromColorCache(t *testing.T) {
	t.Run("remove empty car number to color cache", func(t *testing.T) {
		parkingLotStore := stores.NewParkingLotStoreMock()
		parkingLotService := NewParkingLotService(parkingLotStore)
		err := parkingLotService.removeCarNumberFromColorCache("", "white")
		assert.Error(t, err)
	})
	t.Run("remove car number to empty color cache", func(t *testing.T) {
		parkingLotStore := stores.NewParkingLotStoreMock()
		parkingLotService := NewParkingLotService(parkingLotStore)
		err := parkingLotService.removeCarNumberFromColorCache("KA-01-HH-2708", "")
		assert.Error(t, err)
	})

	t.Run("remove car number to exist color cache", func(t *testing.T) {
		parkingLotStore := stores.NewParkingLotStoreMock()
		parkingLotService := NewParkingLotService(parkingLotStore)
		parkingLotService.carNumbersByColor = map[string]map[string]struct{}{
			"white": {
				"KA-01-HH-2708": struct{}{},
			},
		}
		err := parkingLotService.removeCarNumberFromColorCache("KA-01-HH-2708", "white")
		expected := map[string]map[string]struct{}{
			"white": {},
		}
		assert.Equal(t, expected, parkingLotService.carNumbersByColor)
		assert.Nil(t, err)
	})

	t.Run("add car number to non-exist color cache", func(t *testing.T) {
		parkingLotStore := stores.NewParkingLotStoreMock()
		parkingLotService := NewParkingLotService(parkingLotStore)
		parkingLotService.carNumbersByColor = map[string]map[string]struct{}{
			"white": {},
		}
		err := parkingLotService.removeCarNumberFromColorCache("KA-01-HH-2709", "white")
		expected := map[string]map[string]struct{}{
			"white": {},
		}
		assert.Equal(t, expected, parkingLotService.carNumbersByColor)
		assert.Nil(t, err)
	})
}

func TestService_addSlotToColorCache(t *testing.T) {
	t.Run("add empty slot number to color cache", func(t *testing.T) {
		parkingLotStore := stores.NewParkingLotStoreMock()
		parkingLotService := NewParkingLotService(parkingLotStore)
		err := parkingLotService.addSlotToColorCache(0, "white")
		assert.Error(t, err)
	})
	t.Run("add slot number to empty color cache", func(t *testing.T) {
		parkingLotStore := stores.NewParkingLotStoreMock()
		parkingLotService := NewParkingLotService(parkingLotStore)
		err := parkingLotService.addSlotToColorCache(3, "")
		assert.Error(t, err)
	})

	t.Run("add slot number to exist color cache", func(t *testing.T) {
		parkingLotStore := stores.NewParkingLotStoreMock()
		parkingLotService := NewParkingLotService(parkingLotStore)
		parkingLotService.slotNumbersByColor = map[string]map[int]struct{}{
			"white": {},
		}
		err := parkingLotService.addSlotToColorCache(3, "white")
		expected := map[string]map[int]struct{}{
			"white": {
				3: struct{}{},
			},
		}
		assert.Equal(t, expected, parkingLotService.slotNumbersByColor)
		assert.Nil(t, err)
	})

	t.Run("add slot number to non-exist color cache", func(t *testing.T) {
		parkingLotStore := stores.NewParkingLotStoreMock()
		parkingLotService := NewParkingLotService(parkingLotStore)
		err := parkingLotService.addSlotToColorCache(4, "white")
		expected := map[string]map[int]struct{}{
			"white": {
				4: struct{}{},
			},
		}
		assert.Equal(t, expected, parkingLotService.slotNumbersByColor)
		assert.Nil(t, err)
	})
}

func TestService_removeSlotFromColorCache(t *testing.T) {
	t.Run("remove empty slot number to color cache", func(t *testing.T) {
		parkingLotStore := stores.NewParkingLotStoreMock()
		parkingLotService := NewParkingLotService(parkingLotStore)
		err := parkingLotService.removeSlotFromColorCache(0, "white")
		assert.Error(t, err)
	})
	t.Run("remove slot number to empty color cache", func(t *testing.T) {
		parkingLotStore := stores.NewParkingLotStoreMock()
		parkingLotService := NewParkingLotService(parkingLotStore)
		err := parkingLotService.removeSlotFromColorCache(3, "")
		assert.Error(t, err)
	})

	t.Run("add slot number to exist color cache", func(t *testing.T) {
		parkingLotStore := stores.NewParkingLotStoreMock()
		parkingLotService := NewParkingLotService(parkingLotStore)
		parkingLotService.slotNumbersByColor = map[string]map[int]struct{}{
			"white": {
				3: struct{}{},
			},
		}
		err := parkingLotService.removeSlotFromColorCache(3, "white")
		expected := map[string]map[int]struct{}{
			"white": {},
		}
		assert.Equal(t, expected, parkingLotService.slotNumbersByColor)
		assert.Nil(t, err)
	})

	t.Run("remove slot number to non-exist color cache", func(t *testing.T) {
		parkingLotStore := stores.NewParkingLotStoreMock()
		parkingLotService := NewParkingLotService(parkingLotStore)
		parkingLotService.slotNumbersByColor = map[string]map[int]struct{}{
			"white": {},
		}
		err := parkingLotService.removeSlotFromColorCache(4, "white")
		expected := map[string]map[int]struct{}{
			"white": {},
		}
		assert.Equal(t, expected, parkingLotService.slotNumbersByColor)
		assert.Nil(t, err)
	})
}

func TestService_addSlotToCarNumberCache(t *testing.T) {
	t.Run("add empty slot number to car number cache", func(t *testing.T) {
		parkingLotStore := stores.NewParkingLotStoreMock()
		parkingLotService := NewParkingLotService(parkingLotStore)
		err := parkingLotService.addSlotToCarNumberCache(0, "KA-01-HH-2708")
		assert.Error(t, err)
	})
	t.Run("add slot number to empty car number cache", func(t *testing.T) {
		parkingLotStore := stores.NewParkingLotStoreMock()
		parkingLotService := NewParkingLotService(parkingLotStore)
		err := parkingLotService.addSlotToCarNumberCache(3, "")
		assert.Error(t, err)
	})

	t.Run("add slot number to exist car number cache", func(t *testing.T) {
		parkingLotStore := stores.NewParkingLotStoreMock()
		parkingLotService := NewParkingLotService(parkingLotStore)
		parkingLotService.slotNumberByCarNumber = map[string]int{
			"KA-01-HH-2708": 3,
		}
		err := parkingLotService.addSlotToCarNumberCache(4, "KA-01-HH-2708")
		expected := map[string]int{
			"KA-01-HH-2708": 4,
		}
		assert.Equal(t, expected, parkingLotService.slotNumberByCarNumber)
		assert.Nil(t, err)
	})

	t.Run("add slot number to non-exist car number cache", func(t *testing.T) {
		parkingLotStore := stores.NewParkingLotStoreMock()
		parkingLotService := NewParkingLotService(parkingLotStore)
		err := parkingLotService.addSlotToCarNumberCache(3, "KA-01-HH-2708")
		expected := map[string]int{
			"KA-01-HH-2708": 3,
		}
		assert.Equal(t, expected, parkingLotService.slotNumberByCarNumber)
		assert.Nil(t, err)
	})
}

func TestService_removeSlotFromCarNumberCache(t *testing.T) {
	t.Run("remove slot number to empty car number cache", func(t *testing.T) {
		parkingLotStore := stores.NewParkingLotStoreMock()
		parkingLotService := NewParkingLotService(parkingLotStore)
		err := parkingLotService.removeSlotFromCarNumberCache("")
		assert.Error(t, err)
	})

	t.Run("add slot number to exist car number cache", func(t *testing.T) {
		parkingLotStore := stores.NewParkingLotStoreMock()
		parkingLotService := NewParkingLotService(parkingLotStore)
		parkingLotService.slotNumberByCarNumber = map[string]int{
			"KA-01-HH-2708": 3,
		}
		err := parkingLotService.removeSlotFromCarNumberCache("KA-01-HH-2708")
		expected := map[string]int{}
		assert.Equal(t, expected, parkingLotService.slotNumberByCarNumber)
		assert.Nil(t, err)
	})

	t.Run("add slot number to non-exist car number cache", func(t *testing.T) {
		parkingLotStore := stores.NewParkingLotStoreMock()
		parkingLotService := NewParkingLotService(parkingLotStore)
		err := parkingLotService.removeSlotFromCarNumberCache("KA-01-HH-2708")
		expected := map[string]int{}
		assert.Equal(t, expected, parkingLotService.slotNumberByCarNumber)
		assert.Nil(t, err)
	})
}
