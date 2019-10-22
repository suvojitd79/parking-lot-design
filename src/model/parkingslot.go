package model

//ParkingSlot - conatins the information about a specific spot of the parking lot
type ParkingSlot struct {
	id          uint64  // slot id - [+ve number from 1....N]
	isAvailable bool    // specify whether the spot is available
	vehicle     Vehicle // contains the info of the parked vehicle
}

//ID - getter for id
func (slot ParkingSlot) ID() uint64 {
	return slot.id
}

//IsAvailable - getter for isAvailable
func (slot ParkingSlot) IsAvailable() bool {
	return slot.isAvailable
}

// GetVehicle  - getter for vehicle
func (slot ParkingSlot) GetVehicle() Vehicle {
	return slot.vehicle
}

//SetAvailable - set the availability (removes/adds car)
func (slot *ParkingSlot) SetAvailable(value bool) {
	slot.isAvailable = value
}
