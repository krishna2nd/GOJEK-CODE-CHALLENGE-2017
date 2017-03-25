// Copyright 2017 Krishna Kumar. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package slot implements a simple slot object for vehicles.
// It defines a type Slot with following methods.
// SetNumber(), IsValid(), Allocate(), ISFree(), Free(), GetVehicle()

package slot

import (
	. "ptypes"
	"vehicle"
	"perror"
)

// Slot lower boundry defined as a constant.
const SlotNumberLowerLimit = 1

// Slot defines a Number and a Vehicle.
// If vehicle object is allocated then slot is used
type Slot struct {
	Number  Index
	Vehicle *vehicle.Vehicle
}

// Package based New Object creation function
//  @return:
//		Slot: *Object
func New() *Slot {
	return new(Slot).init()
}

// initialise Object with default values
//  @return:
//		Slot: *Object
func (this *Slot) init() *Slot {
	this.Number = SlotNumberLowerLimit - 1
	this.Vehicle = nil
	return this
}

// Set Slot number to slot object
//  @params:
//      number: string
//  @return:
//		err: error
//		Slot: *Object

func (this *Slot) SetNumber(number Index) (error, *Slot) {
	if number < SlotNumberLowerLimit {
		return perror.ErrSlotNumberInvalid, this
	}
	this.Number = number
	return nil, this
}

// Help to check the slot is valid or not
// Mainly check slot number allocated or not
//  @return:
//		err: bool

func (this *Slot) IsValid() bool {
	return this.Number >= SlotNumberLowerLimit
}

// Set a vehicle object to slot, so that slow will be used
// Slot without valid slot number should show error
// Already using slot should not reused until slot is free
//  @params:
//      vehicle: Vehicle
//  @return:
//		err: error
//		Slot: *Object

func (this *Slot) Allocate(vehicle *vehicle.Vehicle) (error, *Slot) {
	if !this.IsValid() {
		return perror.ErrVehicleAssignInvalidSlot, this
	}
	
	if nil != this.Vehicle {
		return perror.ErrSlotAlreadyAllocated, this
	}
	this.Vehicle = vehicle
	return nil, this
}

// Get vehicle object from allocated slot.
//  @return:
//		err: error
//		vehicle: *Vehicle

func (this *Slot) GetVehicle() *vehicle.Vehicle {
	return this.Vehicle
}

// Remove vehicle object from slot
//  @return:
//		Slot: *Object

func (this *Slot) Free() *Slot {
	this.Vehicle = nil
	return this
}

// Verifies that slot is free or not, if no vehicle allocated
// then vehicle property will be nil
//  @return:
//		isFree: bool

func (this *Slot) IsFree() bool {
	return this.Vehicle == nil
}

