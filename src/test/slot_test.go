package main

import (
	"parking_lot/src/model"
	"testing"
)

func TestSlot(t *testing.T) {

	sampleInputData := []*model.Slot{}
	sampleInputData = append(sampleInputData, model.CreateSlot(uint64(1), false, model.Vehicle{"KKIL-9090", "White"}))

	sampleInputData = append(sampleInputData, model.CreateSlot(uint64(91), false, model.Vehicle{"KKIL-9090-9090", "Black"}))

	sampleInputData = append(sampleInputData, model.CreateSlot(uint64(10), false, model.Vehicle{"KKIL-9090-ISISI", "Pink"}))

	// checking the ids only
	sampleOutputData := []uint64{1, 91, 10}

	for i := 0; i < len(sampleInputData); i++ {

		if sampleInputData[i].ID() != sampleOutputData[i] {

			t.Error("SLOT TEST " + string(i) + " FAILED")

		}

	}

}
