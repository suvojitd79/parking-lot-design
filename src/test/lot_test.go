package main

import (
	"parking_lot/src/model"
	"reflect"
	"testing"
)

func TestParkingLot(t *testing.T) {

	sampleInputData := []*model.ParkingLot{}
	sampleInputData = append(sampleInputData, model.CreateParkingLot(3))

	var str string
	var err error

	str, err = sampleInputData[0].Park(model.Vehicle{RgNumber: "AASS-OP", Col: "Pink"})

	if err != nil {
		t.Error("Test For TestParkingLot Failed")
	} else {
		t.Log(str)
	}

	str, err = sampleInputData[0].Park(model.Vehicle{RgNumber: "BBST-OP", Col: "White"})

	if err != nil {
		t.Error("Test For TestParkingLot Failed")
	} else {
		t.Log(str)
	}

	var data []string

	//
	data, _ = sampleInputData[0].GetRegistrationNumbes("White")
	if !reflect.DeepEqual(data, []string{"BBST-OP"}) {
		t.Error("Test For TestParkingLot Failed > GetRegistrationNumbes")
	} else {
		t.Log(data)
	}

	var dt []uint64

	dt, _ = sampleInputData[0].GetSlotNumbers("White")
	//
	if !reflect.DeepEqual(dt, []uint64{uint64(2)}) {
		t.Error("Test For TestParkingLot Failed > GetSlotNumbers")
	} else {
		t.Log(dt)
	}

	var slotNo uint64

	slotNo, _ = sampleInputData[0].GetSlotNumber("AASS-OP")

	if slotNo != 1 {
		t.Error("Test For TestParkingLot Failed > GetSlotNumber")
	} else {
		t.Log(slotNo)
	}

}
