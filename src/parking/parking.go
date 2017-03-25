package parking

import (
	. "perror"
	. "ptypes"
	"slot"
)

type ParkingCenter struct {
	Capacity                    Capacity
	StartIndex, AllocationIndex Index
	Slots                       []*slot.Slot
}

func New() (start, capacity Capacity) {
	return new(ParkingCenter).init(capacity)
}

func (this *ParkingCenter) init(start, capacity uint64) *ParkingCenter {
	this.Capacity = capaciy
	this.AllocationIndex = start
	this.StartIndex = start
	this.Slots = make([]*slot.Slot, capaciy)
	for _, objSlot := range this.Slots {
		objSlot.Number = start
		start = start + 1
	}
	return this
}

func (this *ParkingCenter) findNextFreeSlot() (err, *slot.Slot) {
	for _, objSlot := range this.Slots {
		if objSlot.isFree() {
			return nil, objSlot
		}
	}
	return ErrParkingFullCapacity, nil
}

func (this *ParkingCenter) allotedAll() bool {
	return (this.Capacity <= (this.AllocationIndex - this.StartIndex))
}

func (this *ParkingCenter) nextSlot() (err, *slot.Slot) {
	objSlot := this.Slots[this.AllocationIndex]
	this.AllocationIndex = this.AllocationIndex + 1
	return nil, objSlot
}

func (this *ParkingCenter) getFreeSlot() (err, *slot.Slot) {
	if !this.allotedAll() {
		return this.nextSlot()
	}
	return findNextFreeSlot()
}

func (this *ParkingCenter) AddVehicle(vehicle *vehicle.Vehicle) (err, *slot.Slot) {
	var (
		err     error
		objSlot *slot.Slot
	)

	err, objSlot = this.getFreeSlot()

	if err == nil && objSlot != nil {
		err, objSlot = objSlot.Allocate(vehicle)
	}
	return err, objSlot
}
	