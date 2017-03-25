package parking

import (
	. "perror"
	. "ptypes"
	. "strconv"
	"testing"
	"vehicle"
)

const (
	Nil     = "Should be nil"
	NoNil   = "Should not be nil"
	NoErr   = "Should not be error"
	NoEmpty = "Should not be empty"
	Full    = "Should be Full capacity"
	NoOrder = "Order Should not changed"
)

func testVehicle(index Index) *vehicle.Vehicle {
	strIndex := FormatUint(uint64(index), 10)
	return vehicle.New("KL "+strIndex, "Color "+strIndex)
}
func addCheck(index Index, pC *ParkingCenter, t *testing.T) {
	err, oSlot := pC.AddVehicle(testVehicle(index))
	if nil != err {
		t.Error(Nil)
	}
	if nil == oSlot {
		t.Error(NoNil)
	}
}

func removeNoErr(index Index, pC *ParkingCenter, t *testing.T) {
	strIndex := FormatUint(uint64(index), 10)
	err, oSlots := pC.RemoveVehicleByNumber("KL " + strIndex)
	if nil != err {
		t.Error(NoErr)
	}

	if 0 == len(oSlots) {
		t.Error(NoEmpty)
	}
}

func addOverFlowCheck(index Index, pC *ParkingCenter, t *testing.T) {
	err, oSlot := pC.AddVehicle(testVehicle(index))
	if nil != oSlot {
		t.Error(Nil)
	}
	if ErrParkingFullCapacity != err {
		t.Error(Full)
	}
}

func TestParkingCenter_Functional(t *testing.T) {
	var (
		start           Index    = 1
		index, capacity Capacity = 0, 100
		pC                       = New(start, capacity)
	)
	for ; uint64(index) < uint64(capacity); index++ {
		addCheck(Index(index), pC, t)
	}
	addOverFlowCheck(Index(index), pC, t)

	removeNoErr(Index(10), pC, t)
	addCheck(Index(10), pC, t)

	for index = 0; uint64(index) < uint64(capacity); index++ {
		if pC.slots[index].Vehicle.Number !=
			testVehicle(Index(index)).Number {
			t.Error(NoOrder)
		}
	}
	for index = 0; uint64(index) < uint64(capacity); index++ {
		removeNoErr(Index(index), pC, t)
	}

}

func BenchmarkParkingCenter_AddVehicle(b *testing.B) {
	var (
		start           Index    = 10000
		index, capacity Capacity = 0, 10000000000000
		pC                       = New(start, capacity)
	)
	for ; uint64(index) < uint64(capacity); index++ {
		pC.AddVehicle(testVehicle(Index(index)))
	}
}

func BenchmarkParkingCenter_AddAndRemove(b *testing.B) {
	var (
		start           Index    = 10000
		index, capacity Capacity = 0, 100000000000
		pC                       = New(start, capacity)
	)
	for ; uint64(index) < uint64(capacity); index++ {
		pC.AddVehicle(testVehicle(Index(index)))
	}
	for ; uint64(index) < uint64(capacity); index++ {
		pC.RemoveVehicleByNumber(testVehicle(Index(index)).Number)
	}
}
