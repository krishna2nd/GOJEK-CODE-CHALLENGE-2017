// Copyright 2017 Krishna Kumar. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package parking manages the complete logic of a parking center
// Create, Add Vehicle, Remove vehicle, Search, Status Report etc.
package parking

import (
	"perror"
	"ptypes"
	"slot"
	"vehicle"
)

// Center object holds the config of all required properties
type Center struct {
	Capacity ptypes.Capacity
	Counter,
	startIndex,
	allocationIndex ptypes.Index
	slots []*slot.Slot
}

// New parking center instance
func New(start ptypes.Index, capacity ptypes.Capacity) *Center {
	return new(Center).
		init(start, capacity)
}

// init parking center instance
func (pc *Center) init(
	start ptypes.Index,
	capacity ptypes.Capacity) *Center {
	pc.Capacity = capacity
	pc.allocationIndex = 0
	pc.startIndex = start
	pc.slots = make([]*slot.Slot, uint64(capacity))
	for idx := range pc.slots {
		pc.slots[idx] = slot.New()
		pc.slots[idx].SetNumber(start)
		start = start + 1
	}
	return pc
}

// findNextFreeSlot to get next free slot
func (pc *Center) findNextFreeSlot() (*slot.Slot, error) {
	for _, objSlot := range pc.slots {
		if objSlot.IsFree() {
			return objSlot, nil
		}
	}
	return nil, perror.ErrParkingFullCapacity
}

// allotedAll to check all slots are allocated at least once or not
func (pc *Center) allotedAll() bool {
	return !(uint64(pc.Capacity) >= uint64(pc.allocationIndex)+1)
}

// nextSlot get next free slot
func (pc *Center) nextSlot() (*slot.Slot, error) {
	objSlot := pc.slots[pc.allocationIndex]
	pc.allocationIndex = pc.allocationIndex + 1
	return objSlot, nil
}

// getFreeSlot get next serial slot or free in between
func (pc *Center) getFreeSlot() (*slot.Slot, error) {
	if !pc.allotedAll() {
		return pc.nextSlot()
	}
	return pc.findNextFreeSlot()
}

// getAllFreeSlot get all free slots serial and in between
func (pc *Center) getAllFreeSlot() []*slot.Slot {
	freeSlots := make([]*slot.Slot, 0)
	for _, s := range pc.slots {
		if s.IsFree() {
			freeSlots = append(freeSlots, s)
		}
	}
	return freeSlots
}

// AddVehicle add vehicle to parking center
func (pc *Center) AddVehicle(vehicle *vehicle.Vehicle) (*slot.Slot, error) {
	var (
		err     error
		objSlot *slot.Slot
	)

	objSlot, err = pc.getFreeSlot()
	if err == nil && objSlot != nil {
		objSlot, err = objSlot.Allocate(vehicle)
		if err == nil {
			pc.Counter = pc.Counter + 1
		}
	}
	return objSlot, err
}

// remove remove vehicle from center and decrement counter
func (pc *Center) remove(s *slot.Slot) {
	s.Free()
	pc.Counter = pc.Counter - 1
}

// RemoveVehicle remove vehicle from center slot list by vehicle object
func (pc *Center) RemoveVehicle(vehicle *vehicle.Vehicle) ([]*slot.Slot, error) {
	var arrSlots = make([]*slot.Slot, 0)
	for _, s := range pc.slots {
		v := s.GetVehicle()
		if nil != v && v.IsEquals(vehicle) {
			pc.remove(s)
			arrSlots = append(arrSlots, s)
		}
	}
	return arrSlots, nil
}

// RemoveVehicleByNumber remove vehicle from center slot list by vehicle number
func (pc *Center) RemoveVehicleByNumber(number string) ([]*slot.Slot, error) {
	slots, err := pc.GetSlotsBy("number", number)
	for _, v := range slots {
		pc.remove(v)
	}
	return slots, err
}

// RemoveVehicleBySlotNumber  remove slot from center slot list by slot number
func (pc *Center) RemoveVehicleBySlotNumber(Number ptypes.Index) (*slot.Slot, error) {
	oSlot, err := pc.GetSlot(Number)
	if nil != err {
		return oSlot, err
	}

	pc.remove(oSlot)
	return oSlot, nil
}

// GetSlot from center slot list by vehicle number
func (pc *Center) GetSlot(Number ptypes.Index) (*slot.Slot, error) {
	if (Number) < (ptypes.Index(pc.Capacity)+pc.startIndex) &&
		Number >= pc.startIndex {
		return pc.slots[Number-pc.startIndex], nil
	}
	return nil, perror.ErrInvalidGetSlotNumber
}

// GetSlotsBy vehicle property { number, color }
func (pc *Center) GetSlotsBy(property, value string) ([]*slot.Slot, error) {
	var arrSlots = make([]*slot.Slot, 0)
	var val string
	for _, s := range pc.slots {
		v := s.GetVehicle()
		if nil != v {
			switch property {
			case "number":
				val = v.GetNumber()
				break
			case "color":
				val = v.GetColour()
				break
			}
			if value == val {
				arrSlots = append(arrSlots, s)
			}
		}
	}
	return arrSlots, nil
}
