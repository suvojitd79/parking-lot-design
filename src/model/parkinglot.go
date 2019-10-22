package model

//ParkingLot -
type ParkingLot struct {
	size                        uint64 // specify the size of the parking lot (N)
	slots                       []Slot
	isAvailable                 bool                // whether the lot is full or not
	colorRegistrationNumbesDict map[string][]string // returns a list of registrations on passing a color as a key (color - > reg ids)
	colorSlotsDict              map[string][]uint64 // returns the slot ids for a specific color (color -> slots)
	registrationNumbeSlotDict   map[string]string   // returns the slot number for a specific registration number
}

//CreateParkingLot - works as a constructor for the ParkingLot
func CreateParkingLot(size uint64) *ParkingLot {

	parkingLot := ParkingLot{}
	parkingLot.size = size
	parkingLot.slots = []Slot{}
	parkingLot.isAvailable = (size > 0)
	parkingLot.colorRegistrationNumbesDict = map[string][]string{}
	parkingLot.colorSlotsDict = map[string][]uint64{}
	parkingLot.registrationNumbeSlotDict = map[string]string{}
	return &parkingLot
}
