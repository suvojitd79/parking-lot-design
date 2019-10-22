package model

import "errors"

//ParkingLot -
type ParkingLot struct {
	capacity                    uint64 // specify the size of the parking lot (N)
	availableSpots              uint64 // specify the number of available spots
	slots                       []Slot
	colorRegistrationNumbesDict map[string][]string // returns a list of registrations on passing a color as a key (color - > reg ids)
	colorSlotNumbersDict        map[string][]uint64 // returns the slot ids for a specific color (color -> slots)
	registrationNumbeSlotDict   map[string]uint64   // returns the slot number for a specific registration number
}

//CreateParkingLot - works as a constructor for the ParkingLot
func CreateParkingLot(size uint64) *ParkingLot {

	parkingLot := ParkingLot{}
	parkingLot.capacity = size
	parkingLot.availableSpots = size
	parkingLot.slots = []Slot{}
	parkingLot.colorRegistrationNumbesDict = map[string][]string{}
	parkingLot.colorSlotNumbersDict = map[string][]uint64{}
	parkingLot.registrationNumbeSlotDict = map[string]uint64{}
	return &parkingLot
}

//GetCapacity - returns the capacity of the parking lot
func (lot ParkingLot) GetCapacity() uint64 {
	return lot.capacity
}

//GetAvailableSpots - return the total number of available spots
func (lot ParkingLot) GetAvailableSpots() uint64 {
	return lot.availableSpots
}

//IsParkingAvailable - check whether parking space available
func (lot ParkingLot) IsParkingAvailable() bool {
	return lot.availableSpots > 0
}

//GetSlots - returns all the slots
func (lot ParkingLot) GetSlots() []Slot {
	return lot.slots
}

//GetRegistrationNumbes - gives the list of reg. no associated with a color
func (lot ParkingLot) GetRegistrationNumbes(color string) ([]string, error) {

	if val, found := lot.colorRegistrationNumbesDict[color]; found {
		return val, nil
	}
	return []string{}, errors.New("Not found") // empty list if not found
}

//GetSlotNumbers - gives the list of reg. no associated with a color
func (lot ParkingLot) GetSlotNumbers(color string) ([]uint64, error) {

	if val, found := lot.colorSlotNumbersDict[color]; found {
		return val, nil
	}
	return []uint64{}, errors.New("Not found") // empty list if not found
}

//GetSlotNumber - gives the list of reg. no associated with a registration number
func (lot ParkingLot) GetSlotNumber(registration string) (uint64, error) {

	if val, found := lot.registrationNumbeSlotDict[registration]; found {
		return val, nil
	}
	return 0, errors.New("Not found") // empty list if not found
}
