package model

//Slot - conatins the information about a specific spot of the parking lot
type Slot struct {
	id          uint64  // slot id - [+ve number from 1....N]
	isAvailable bool    // specify whether the spot is available
	vehicle     Vehicle // contains the info of the parked vehicle
}

//ID - getter for id
func (slot Slot) ID() uint64 {
	return slot.id
}

//IsAvailable - getter for isAvailable
func (slot Slot) IsAvailable() bool {
	return slot.isAvailable
}

// GetVehicle  - getter for vehicle
func (slot Slot) GetVehicle() Vehicle {
	return slot.vehicle
}

//SetAvailable - set the availability (removes/adds car)
func (slot *Slot) SetAvailable(value bool) {
	slot.isAvailable = value
}
