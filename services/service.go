package services

import (
	"errors"
	"fmt"
	"parking_lot/models"
	"parking_lot/stores"
	"strconv"
	"strings"
	"unicode"
)

type IParkingLotService interface {
	CreateParkingLot(slotAmount int) ([]*models.ParkingSlot, error)
	Park(carNumber string, carColor string) (int, error)
	Leave(slotNumber int) error
	GetStatus() (string, error)
	GetParkedCarNumbersByColor(carColor string) (string, error)
	GetParkedSlotNumbersByColor(carColor string) (string, error)
	GetParkedSlotNumberByCarNumber(carNumber string) (int, error)
}

type ParkingLotService struct {
	store stores.IParkingLotStore

	// Local cache
	carNumbersByColor     map[string]map[string]struct{}
	slotNumbersByColor    map[string]map[int]struct{}
	slotNumberByCarNumber map[string]int
}

func NewParkingLotService(parkingLotStore stores.IParkingLotStore) *ParkingLotService {
	return &ParkingLotService{
		store:                 parkingLotStore,
		carNumbersByColor:     map[string]map[string]struct{}{},
		slotNumbersByColor:    map[string]map[int]struct{}{},
		slotNumberByCarNumber: map[string]int{},
	}
}

func (svc *ParkingLotService) CreateParkingLot(slotAmount int) ([]*models.ParkingSlot, error) {
	parkingSlots, err := svc.store.CreateParkingLot(slotAmount)
	if err != nil {
		return nil, err
	}

	return parkingSlots, nil
}

func (svc *ParkingLotService) Park(carNumber string, carColor string) (int, error) {
	if len(carNumber) == 0 {
		return 0, errors.New("car's number is empty")
	}
	if len(carColor) == 0 {
		return 0, errors.New("car's color is empty")
	}

	parkingSlots, err := svc.store.GetParkingSlots()
	if err != nil {
		return 0, err
	}

	carColor = strings.ToLower(carColor)
	for idx, slot := range parkingSlots {
		// Skip if the current lot is occupied.
		if slot.ParkedCar != nil {
			continue
		}

		// Update parking lot data.
		parkedCar := &models.Car{}
		parkedCar.RegistrationNumber = carNumber
		parkedCar.Color = carColor
		_, err := svc.store.UpdateParkingLot(idx, parkedCar)
		if err != nil {
			return 0, err
		}

		// Update local cache.
		slotNumber := idx + 1
		svc.addCarNumberToColorCache(carNumber, carColor)
		svc.addSlotToColorCache(slotNumber, carColor)
		svc.addSlotToCarNumberCache(slotNumber, carNumber)

		return slotNumber, nil
	}

	return 0, errors.New("parking lot is full")
}

func (svc *ParkingLotService) Leave(slotNumber int) error {
	parkingSlot, err := svc.store.GetParkingSlotDetail(slotNumber - 1)
	if err != nil {
		return err
	}

	if parkingSlot == nil {
		return errors.New("this slot is not available")
	}
	if parkingSlot.ParkedCar == nil {
		return errors.New("this slot is not occupied")
	}

	_, err = svc.store.UpdateParkingLot(slotNumber-1, nil)
	if err != nil {
		return err
	}

	carColor := parkingSlot.ParkedCar.Color
	carNumber := parkingSlot.ParkedCar.RegistrationNumber
	svc.removeCarNumberFromColorCache(carNumber, carColor)
	svc.removeSlotFromColorCache(slotNumber, carColor)
	svc.removeSlotFromCarNumberCache(carNumber)

	return nil
}

func (svc *ParkingLotService) GetStatus() (string, error) {
	parkingSlots, err := svc.store.GetParkingSlots()
	if err != nil {
		return "", err
	}

	status := ""
	for idx, slot := range parkingSlots {
		if slot.ParkedCar == nil {
			continue
		}
		r := []rune(slot.ParkedCar.Color)
		r[0] = unicode.ToUpper(r[0])
		formatColor := string(r)
		status += fmt.Sprintf("%d\t\t%s\t\t%s\n", idx+1, slot.ParkedCar.RegistrationNumber, formatColor)
	}

	return status, nil
}

func (svc *ParkingLotService) GetParkedCarNumbersByColor(carColor string) (string, error) {
	if len(carColor) == 0 {
		return "", errors.New("car's color is invalid")
	}

	cData, ok := svc.carNumbersByColor[carColor]
	if !ok || len(cData) == 0 {
		return "", errors.New("not found")
	}

	results := make([]string, 0, len(cData))
	for carNumber := range cData {
		if len(carNumber) == 0 {
			continue
		}
		results = append(results, carNumber)
	}

	return strings.Join(results, ", "), nil
}

func (svc *ParkingLotService) GetParkedSlotNumbersByColor(carColor string) (string, error) {
	if len(carColor) == 0 {
		return "", errors.New("car's color is invalid")
	}

	cData, ok := svc.slotNumbersByColor[carColor]
	if !ok || len(cData) == 0 {
		return "", errors.New("not found")
	}

	results := make([]string, 0, len(cData))
	for slotNumber := range cData {
		results = append(results, strconv.Itoa(slotNumber+1))
	}

	return strings.Join(results, ", "), nil
}

func (svc *ParkingLotService) GetParkedSlotNumberByCarNumber(carNumber string) (int, error) {
	if len(carNumber) == 0 {
		return 0, errors.New("car's number is invalid")
	}

	slotNumber, ok := svc.slotNumberByCarNumber[carNumber]
	if !ok {
		return 0, errors.New("not found")
	}

	return slotNumber + 1, nil
}

func (svc *ParkingLotService) addCarNumberToColorCache(carNumber string, carColor string) error {
	if len(carNumber) == 0 {
		return errors.New("car's number is empty")
	}
	if len(carColor) == 0 {
		return errors.New("car's color is empty")
	}

	cData, ok := svc.carNumbersByColor[carColor]
	if !ok {
		svc.carNumbersByColor[carColor] = map[string]struct{}{
			carNumber: {},
		}
	} else {
		cData[carNumber] = struct{}{}
		svc.carNumbersByColor[carColor] = cData
	}

	return nil
}

func (svc *ParkingLotService) removeCarNumberFromColorCache(carNumber string, carColor string) error {
	if len(carNumber) == 0 {
		return errors.New("car's number is empty")
	}
	if len(carColor) == 0 {
		return errors.New("car's color is empty")
	}

	if cData, ok := svc.carNumbersByColor[carColor]; ok {
		delete(cData, carNumber)
		svc.carNumbersByColor[carColor] = cData
	}

	return nil
}

func (svc *ParkingLotService) addSlotToColorCache(slotNumber int, carColor string) error {
	if slotNumber < 1 {
		return errors.New("slot number is invalid")
	}
	if len(carColor) == 0 {
		return errors.New("car's color is empty")
	}

	cData, ok := svc.slotNumbersByColor[carColor]
	if !ok {
		svc.slotNumbersByColor[carColor] = map[int]struct{}{
			slotNumber: {},
		}
	} else {
		cData[slotNumber] = struct{}{}
		svc.slotNumbersByColor[carColor] = cData
	}

	return nil
}

func (svc *ParkingLotService) removeSlotFromColorCache(slotNumber int, carColor string) error {
	if slotNumber < 1 {
		return errors.New("slot number is invalid")
	}
	if len(carColor) == 0 {
		return errors.New("car's color is empty")
	}

	if cData, ok := svc.slotNumbersByColor[carColor]; ok {
		delete(cData, slotNumber)
		svc.slotNumbersByColor[carColor] = cData
	}

	return nil
}

func (svc *ParkingLotService) addSlotToCarNumberCache(slotNumber int, carNumber string) error {
	if slotNumber < 1 {
		return errors.New("slot number is invalid")
	}
	if len(carNumber) == 0 {
		return errors.New("car's number is empty")
	}

	svc.slotNumberByCarNumber[carNumber] = slotNumber
	return nil
}

func (svc *ParkingLotService) removeSlotFromCarNumberCache(carNumber string) error {
	if len(carNumber) == 0 {
		return errors.New("car's number is empty")
	}

	delete(svc.slotNumberByCarNumber, carNumber)
	return nil
}
