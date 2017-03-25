package slot

import (
	"vehicle"
	"perror"
)

const SlotNumberLowerLimit = 1

type Slot struct {
	Number  uint64
	Vehicle *vehicle.Vehicle
}

func New() *Slot {
	return new(Slot).Init()
}

func (this *Slot) Init() *Slot {
	this.Number = SlotNumberLowerLimit - 1
	this.Vehicle = nil
	return this
}

func (this *Slot) SetNumber(number uint64) (error, *Slot) {
	if number < SlotNumberLowerLimit {
		return perror.ErrSlotNumberInvalid, this
	}
	this.Number = number
	return nil, this
}

func (this *Slot) Valid() bool {
	return this.Number >= SlotNumberLowerLimit
}

func (this *Slot) Allocate(vehicle *vehicle.Vehicle) (error, *Slot) {
	if !this.Valid() {
		return perror.ErrVehicleAssignInvalidSlot, this
	}
	
	if nil != this.Vehicle {
		return perror.ErrSlotAlreadyAllocated, this
	}
	this.Vehicle = vehicle
	return nil, this
}

func (this *Slot) GetVehicle() *vehicle.Vehicle {
	return this.Vehicle
}

func (this *Slot) Free() *Slot {
	this.Vehicle = nil
	return this
}

func (this *Slot) IsFree() bool {
	return this.Vehicle == nil
}

