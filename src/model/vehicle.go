package model

//Vehicle - model of base vehicle class
type Vehicle struct {
	RgNumber string `json:"registrationNumber"`
	Col      string `json:"color"`
}

// RegistrationNumber - returns the registration number of the vehicle
func (vehicle Vehicle) RegistrationNumber() string {
	return vehicle.RgNumber
}

// Color - returns the color of the vehicle
func (vehicle Vehicle) Color() string {
	return vehicle.Col
}
