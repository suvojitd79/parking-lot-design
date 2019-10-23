package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"parking_lot/src/model"
	"testing"
)

type VehicleTest struct {
	Input  model.Vehicle `json:"input"`
	Output model.Vehicle `json:"output"`
}

func TestVehicle(t *testing.T) {

	jsonFile, err := os.Open("./config/vehicle.json")
	if err != nil {
		t.Error("VEHICLE TEST FAILED")
	}
	defer jsonFile.Close()
	byteData, e := ioutil.ReadAll(jsonFile)
	if e != nil {
		t.Error("VEHICLE TEST FAILED")
	}
	var testVehicles []VehicleTest
	json.Unmarshal(byteData, &testVehicles)
	for i := 0; i < len(testVehicles); i++ {
		if testVehicles[i].Input.RegistrationNumber() != testVehicles[i].Output.RegistrationNumber() || testVehicles[i].Input.Color() != testVehicles[i].Output.Color() {
			t.Error("VEHICLE TEST " + string(i+1) + " FAILED")
		}
	}
}
