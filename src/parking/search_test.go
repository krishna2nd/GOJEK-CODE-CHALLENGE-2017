package parking

import (
	. "ptypes"
	"testing"
)

const (
	NoSlot = "Should be one slot free"
	Slot11 = "Slot Number should be 11"

	Color        = "Color 12"
	NoColor12    = "Color 12 should present"
	NoColor12Val = "Colour value should be 'Color 12'"

	NoNumber12    = "Number 12 should present"
	NoNumber12Val = "Number value should be 'KL 12'"
)

var (
	start           Index    = 1
	index, capacity Capacity = 0, 100
	pC                       = New(start, capacity)
	tVehicle                 = testVehicle(Index(12))
)

func AddDataForSearch() {
	for ; uint64(index) < uint64(capacity); index++ {
		pC.AddVehicle(testVehicle(Index(index)))
	}
}
func TestParkingCenter_ReportFreeSlots(t *testing.T) {
	AddDataForSearch()
	removeNoErr(Index(10), pC, t)
	oSlots := pC.ReportFreeSlots()
	if 1 != len(oSlots) {
		t.Error(NoSlot)
	} else if oSlots[0].Number != Index(11) {
		t.Error(Slot11)
	}
}

func TestParkingCenter_ReportVehicleByColor(t *testing.T) {
	AddDataForSearch()
	oSlots, _ := pC.ReportVehicleByColor(tVehicle.GetColour())
	if 1 != len(oSlots) {
		t.Error(NoColor12)
	} else if oSlots[0].Vehicle.Color != tVehicle.GetColour() {
		t.Error(NoColor12Val)
	}
}

func TestParkingCenter_ReportVehicleByNumber(t *testing.T) {
	AddDataForSearch()
	oSlots, _ := pC.ReportVehicleByNumber(tVehicle.GetNumber())
	if 1 != len(oSlots) {
		t.Error(NoNumber12)
	} else if oSlots[0].Vehicle.GetNumber() != tVehicle.GetNumber() {
		t.Error(NoNumber12Val)
	}
}
