package model

import (
	"errors"
	"strconv"
)

//ParkingLot -
type ParkingLot struct {
	capacity                     uint64 // specify the size of the parking lot (N)
	availableSpots               uint64 // specify the number of available spots
	slots                        []Slot
	colorRegistrationNumbersDict map[string][]string // returns a list of registrations on passing a color as a key (color - > reg ids)
	colorSlotNumbersDict         map[string][]uint64 // returns the slot ids for a specific color (color -> slots)
	registrationNumberSlotDict   map[string]uint64   // returns the slot number for a specific registration number
	nextAvailableSpotNumber      []uint64            // holds the info regarding next available spot, first value will always gives the next available spot
}

//CreateParkingLot - works as a constructor for the ParkingLot
func CreateParkingLot(size uint64) *ParkingLot {

	parkingLot := ParkingLot{}
	parkingLot.capacity = size
	parkingLot.availableSpots = size
	parkingLot.slots = []Slot{}
	parkingLot.colorRegistrationNumbersDict = map[string][]string{}
	parkingLot.colorSlotNumbersDict = map[string][]uint64{}
	parkingLot.registrationNumberSlotDict = map[string]uint64{}
	if size > 0 {
		parkingLot.nextAvailableSpotNumber = []uint64{}
		var i uint64
		for i = 1; i <= size; i++ {
			parkingLot.nextAvailableSpotNumber = append(parkingLot.nextAvailableSpotNumber, i)
		}

		// initialize the next available spot as 1 if size > 0
	} else {
		parkingLot.nextAvailableSpotNumber = []uint64{} // not available
	}

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
	return lot.availableSpots > 0 && len(lot.nextAvailableSpotNumber) > 0
}

//GetSlots - returns all the slots
func (lot ParkingLot) GetSlots() []Slot {
	return lot.slots
}

//GetSlotByID - return slot with specific id (1 based)
func (lot ParkingLot) GetSlotByID(id uint64) (Slot, error) {
	if uint64(len(lot.slots)) < id { // not available
		return Slot{}, errors.New("Not exists")
	}
	return lot.slots[id-1], nil
}

//M1 =============================================================================

//GetRegistrationNumbes - gives the list of reg. no associated with a color
func (lot *ParkingLot) GetRegistrationNumbes(color string) ([]string, error) {

	if val, found := lot.colorRegistrationNumbersDict[color]; found {
		return val, nil
	}
	return []string{}, errors.New("Not found") // empty list if not found
}

//RemoveRegistrationNumber - removes reg no. associated with a color
func (lot *ParkingLot) RemoveRegistrationNumber(vh Vehicle) {

	vhColor := vh.Color()
	vhRegistrationNo := vh.RegistrationNumber()

	if val, found := lot.colorRegistrationNumbersDict[vhColor]; found {

		for index, regNo := range val {

			if regNo == vhRegistrationNo {

				lot.colorRegistrationNumbersDict[vhColor] = append(val[:index], val[index+1:]...)
			}

		}
	}
}

//AddRegistrationNumber - add reg no. associated with a color
func (lot *ParkingLot) AddRegistrationNumber(vh Vehicle) {

	lot.RemoveRegistrationNumber(vh) // double check before putting an item in the map
	vhColor := vh.Color()
	vhRegistrationNo := vh.RegistrationNumber()

	if val, found := lot.colorRegistrationNumbersDict[vhColor]; found {
		lot.colorRegistrationNumbersDict[vhColor] = append(val, vhRegistrationNo)
	} else {
		lot.colorRegistrationNumbersDict[vhColor] = []string{vhRegistrationNo}
	}

}

//M2 =============================================================================

//GetSlotNumbers - gives the list of slot. no associated with a color
func (lot *ParkingLot) GetSlotNumbers(color string) ([]uint64, error) {

	if val, found := lot.colorSlotNumbersDict[color]; found {
		return val, nil
	}
	return []uint64{}, errors.New("Not found") // empty list if not found
}

//RemoveSlotNumbers - removes the slot. no associated with a color
func (lot *ParkingLot) RemoveSlotNumbers(vh Vehicle) {

	vhColor := vh.Color()
	vhRegistrationNo := vh.RegistrationNumber()

	slotID, found := lot.registrationNumberSlotDict[vhRegistrationNo]

	if !found {
		return // not found
	}

	if val, f := lot.colorSlotNumbersDict[vhColor]; f {

		for index, slotNo := range val {

			if slotNo == slotID {

				lot.colorSlotNumbersDict[vhColor] = append(val[:index], val[index+1:]...)
			}

		}

	}
}

//AddSlotNumbers - add the slot. no associated with a color
func (lot *ParkingLot) AddSlotNumbers(vh Vehicle, id uint64) {

	lot.RemoveSlotNumbers(vh) // double check before putting an item in the map

	vhColor := vh.Color()

	if val, found := lot.colorSlotNumbersDict[vhColor]; found {
		lot.colorSlotNumbersDict[vhColor] = append(val, id)
	} else {
		lot.colorSlotNumbersDict[vhColor] = []uint64{id}
	}

}

//M3 =============================================================================

//GetSlotNumber - gives the slot number associated with a registration number
func (lot *ParkingLot) GetSlotNumber(registration string) (uint64, error) {

	if val, found := lot.registrationNumberSlotDict[registration]; found {
		return val, nil
	}
	return 0, errors.New("Not found") // empty list if not found
}

//RemoveSlotNumber - removes slot number associated with a registration number
func (lot *ParkingLot) RemoveSlotNumber(vh Vehicle) {
	regNo := vh.RegistrationNumber()
	if _, found := lot.registrationNumberSlotDict[regNo]; found {
		delete(lot.registrationNumberSlotDict, regNo)
	}
}

//AddSlotNumber - add the slot number associated with a registration number
func (lot *ParkingLot) AddSlotNumber(vh Vehicle, slot uint64) {
	regNo := vh.RegistrationNumber()
	lot.RemoveSlotNumber(vh)
	lot.registrationNumberSlotDict[regNo] = slot
}

//========================================================================

//pollAvailableSpotIndex - returns & remove the first item of nextAvailableSpotNumber
func (lot *ParkingLot) pollAvailableSpotIndex() uint64 {

	number := lot.nextAvailableSpotNumber[0]

	lot.nextAvailableSpotNumber = lot.nextAvailableSpotNumber[1:] // remove the first element

	return number
}

//Park - park a vehicle if empty spot is found
func (lot *ParkingLot) Park(vh Vehicle) (string, error) {

	// if the parking lot is full, return error msg
	if !lot.IsParkingAvailable() {
		return "", errors.New("Sorry, parking lot is full")
	}

	availableSpotNumber := lot.pollAvailableSpotIndex()

	// check whether slot exists or not
	_, err := lot.GetSlotByID(availableSpotNumber)

	// slot already exists
	if err == nil {
		lot.slots[availableSpotNumber-1] = Slot{id: availableSpotNumber, isAvailable: false, vehicle: vh}
	} else {
		// insert the vehicle into the new specific spot
		lot.slots = append(lot.slots, Slot{id: availableSpotNumber, isAvailable: false, vehicle: vh})
	}

	//AddRegistrationNumber - add reg no. associated with a color
	lot.AddRegistrationNumber(vh)

	//AddSlotNumber - add the slot number associated with a registration number
	lot.AddSlotNumber(vh, availableSpotNumber)

	//AddSlotNumbers - add the reg. no associated with a color
	lot.AddSlotNumbers(vh, availableSpotNumber)

	return "Allocated slot number: " + strconv.FormatUint(availableSpotNumber, 10), nil
}
