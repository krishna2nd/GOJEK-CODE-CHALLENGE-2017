// Copyright 2017 Krishna Kumar. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package slot implements a simple slot object for vehicles.
// It defines a type Slot with following methods.
// SetNumber(), IsValid(), Allocate(), ISFree(), Free(), GetVehicle()
package slot

import (
	"perror"
	"ptypes"
	"vehicle"
)

// SlotNumberLowerLimit is slot lower bound defined as a constant.
const SlotNumberLowerLimit = 1

// Slot defines a Number and a Vehicle.
// If vehicle object is allocated then slot is used
type Slot struct {
	Number  ptypes.Index
	Vehicle *vehicle.Vehicle
}

// New Slot Object creation function
//  @return:
//		Slot: *Object
func New() *Slot {
	return new(Slot).init()
}

// initialise Object with default values
//  @return:
//		Slot: *Object
func (sl *Slot) init() *Slot {
	sl.Number = SlotNumberLowerLimit - 1
	sl.Vehicle = nil
	return sl
}

// SetNumber Set slot number to slot object
//  @params:
//      number: string
//  @return:
//		Slot: *Object
//		err: error
func (sl *Slot) SetNumber(number ptypes.Index) (*Slot, error) {
	if !IsValidSlotNumber(number) {
		return sl, perror.ErrSlotNumberInvalid
	}
	sl.Number = number
	return sl, nil
}

// GetNumber get slot number from slot object
//  @params: (void)
//      number: string
//  @return:
//		number: Index
func (sl *Slot) GetNumber() ptypes.Index {
	return sl.Number
}

// IsValid help to check the slot is valid or not
// Mainly check slot number allocated or not
//  @return:
//		flag: bool
func (sl *Slot) IsValid() bool {
	return sl.Number >= SlotNumberLowerLimit
}

// IsValidSlotNumber help to check the slot number is valid or not
//  @return:
//		flag : bool
func IsValidSlotNumber(Number ptypes.Index) bool {
	return (Number >= SlotNumberLowerLimit)
}

// Allocate set a vehicle object to slot, so that slow will be used
// Slot without valid slot number should show error
// Already using slot should not reused until slot is free
//  @params:
//      vehicle: Vehicle
//  @return:
//		Slot: *Object
//		err: error
func (sl *Slot) Allocate(vehicle *vehicle.Vehicle) (*Slot, error) {
	if !sl.IsValid() {
		return sl, perror.ErrVehicleAssignInvalidSlot
	}

	if nil != sl.Vehicle {
		return sl, perror.ErrSlotAlreadyAllocated
	}
	sl.Vehicle = vehicle
	return sl, nil
}

// GetVehicle get vehicle object from allocated slot.
//  @return:
//		err: error
//		vehicle: *Vehicle
func (sl *Slot) GetVehicle() *vehicle.Vehicle {
	return sl.Vehicle
}

// Free remove vehicle object from slot
//  @return:
//		Slot: *Object
func (sl *Slot) Free() *Slot {
	sl.Vehicle = nil
	return sl
}

// IsFree Verifies that slot is free or not, if no vehicle allocated
// then vehicle property will be nil
//  @return:
//		isFree: bool
func (sl *Slot) IsFree() bool {
	return sl.Vehicle == nil
}
