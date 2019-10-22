package model

//ParkingLot -
type ParkingLot struct {
	size                        uint64 // specify the size of the parking lot (N)
	slots                       []ParkingSlot
	isAvailable                 bool                // whether the lot is full or not
	colorRegistrationNumbesDict map[string][]string // returns a list of registrations on passing a color as a key (color - > reg ids)
	colorSlotsDict              map[string][]uint64 // returns the slot ids for a specific color (color -> slots)
	registrationNumbeSlotDict   map[string]string   // returns the slot number for a specific registration number
}
