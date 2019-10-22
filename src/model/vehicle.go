package model

//Vehicle - model of base vehicle class
type Vehicle struct {
	registrationNumber string
	color              string
}

// RegistrationNumber - returns the registration number of the vehicle
func (vehicle Vehicle) RegistrationNumber() string {
	return vehicle.registrationNumber
}

//Color - returns the color of the vehicle
func (vehicle Vehicle) Color() string {
	return vehicle.color
}
