package parking

import (
	. "perror"
	. "ptypes"
	"slot"
	"vehicle"
)

type ParkingCenter struct {
	Capacity Capacity
	Counter,
	startIndex,
	allocationIndex Index
	slots []*slot.Slot
}

func New(start Index, capacity Capacity) *ParkingCenter {
	return new(ParkingCenter).
		init(start, capacity)
}

func (this *ParkingCenter) init(start Index, capacity Capacity) *ParkingCenter {
	this.Capacity = capacity
	this.allocationIndex = 0
	this.startIndex = start
	this.slots = make([]*slot.Slot, uint64(capacity))
	for idx, _ := range this.slots {
		this.slots[idx] = slot.New()
		this.slots[idx].SetNumber(start)
		start = start + 1
	}
	return this
}

func (this *ParkingCenter) findNextFreeSlot() (error, *slot.Slot) {
	for _, objSlot := range this.slots {
		if objSlot.IsFree() {
			return nil, objSlot
		}
	}
	return ErrParkingFullCapacity, nil
}

func (this *ParkingCenter) allotedAll() bool {
	return !(uint64(this.Capacity) >= uint64(this.allocationIndex)+1)
}

func (this *ParkingCenter) nextSlot() (error, *slot.Slot) {
	objSlot := this.slots[this.allocationIndex]
	this.allocationIndex = this.allocationIndex + 1
	return nil, objSlot
}

func (this *ParkingCenter) getFreeSlot() (error, *slot.Slot) {
	if !this.allotedAll() {
		return this.nextSlot()
	}
	return this.findNextFreeSlot()
}

func (this *ParkingCenter) getAllFreeSlot() []*slot.Slot {
	freeSlots := make([]*slot.Slot, 0)
	for _, s := range this.slots {
		if s.IsFree() {
			freeSlots = append(freeSlots, s)
		}
	}
	return freeSlots
}

func (this *ParkingCenter) AddVehicle(vehicle *vehicle.Vehicle) (error, *slot.Slot) {
	var (
		err     error
		objSlot *slot.Slot
	)

	err, objSlot = this.getFreeSlot()
	if err == nil && objSlot != nil {
		err, objSlot = objSlot.Allocate(vehicle)
		if err == nil {
			this.Counter = this.Counter + 1
		}
	}
	return err, objSlot
}

func (this *ParkingCenter) remove(s *slot.Slot) {
	s.Free()
	this.Counter = this.Counter - 1
}

func (this *ParkingCenter) RemoveVehicle(vehicle *vehicle.Vehicle) (error, []*slot.Slot) {
	var arrSlots = make([]*slot.Slot, 0)
	for _, s := range this.slots {
		v := s.GetVehicle()
		if nil != v && v.IsEquals(vehicle) {
			this.remove(s)
			arrSlots = append(arrSlots, s)
		}
	}
	return nil, arrSlots
}

func (this *ParkingCenter) RemoveVehicleByNumber(number string) (error, []*slot.Slot) {
	_, oSlots := this.GetSlotsBy("number", number)
	for _, v := range oSlots {
		this.remove(v)
	}
	return nil, oSlots
}

func (this *ParkingCenter) RemoveVehicleBySlotNumber(Number Index) (error, *slot.Slot) {
	err, oSlot := this.GetSlot(Number)
	if nil != err {
		return err, oSlot
	}
	
	this.remove(oSlot)
	
	return nil, oSlot
}

func (this *ParkingCenter) GetSlot(Number Index) (error, *slot.Slot) {
	if (Number) < (Index(this.Capacity) + this.startIndex) &&
		Number >= this.startIndex {
		return nil, this.slots[Number - this.startIndex]
	}
	return ErrInvalidGetSlotNumber, nil
}

func (this *ParkingCenter) GetSlotsBy(property, value string) (error, []*slot.Slot) {
	var arrSlots = make([]*slot.Slot, 0)
	var val string
	for _, s := range this.slots {
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
	return nil, arrSlots
}
