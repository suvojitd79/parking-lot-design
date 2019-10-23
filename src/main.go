package main

import (
	"fmt"
	"parking_lot/src/model"
)

func main() {

	lot, str := model.CreateParkingLot(2)

	fmt.Println(str)

	var status string
	var err error

	status, err = lot.Park(model.Vehicle{RgNumber: "XX90", Col: "Black"})

	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(status)
	}

	status, err = lot.Park(model.Vehicle{RgNumber: "XXPO", Col: "Black"})

	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(status)
	}

	status, err = lot.Park(model.Vehicle{RgNumber: "XX99", Col: "Black"})

	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(status)
	}

	lot.Status()

	status, err = lot.Leave(uint64(10))
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(status)
	}

	lot.Status()

}
