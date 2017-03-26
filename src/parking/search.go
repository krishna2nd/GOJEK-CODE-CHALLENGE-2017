package parking

import (
	. "perror"
	. "ptypes"
	"slot"
)

func (this *ParkingCenter) ReportVehicleByColor(number string) (error, []*slot.Slot) {
	err, oSlots := this.GetSlotsBy("color", number)
	if Zero == len(oSlots) {
		err = ErrNotFound
	}
	return err, oSlots
}

func (this *ParkingCenter) ReportVehicleByNumber(number string) (error, []*slot.Slot) {
	err, oSlots := this.GetSlotsBy("number", number)
	if Zero == len(oSlots) {
		err = ErrNotFound
	}
	return err, oSlots
}

func (this *ParkingCenter) ReportFreeSlots() []*slot.Slot {
	return this.getAllFreeSlot()
}

func (this *ParkingCenter) ReportFilledSlots() (error, []*slot.Slot) {
	allocSlots := make([]*slot.Slot, 0)
	if Index(Zero) == this.Counter {
		return ErrNoFilledSlots, nil
	}
	for _, s := range this.slots {
		if !s.IsFree() {
			allocSlots = append(allocSlots, s)
		}
	}
	return nil, allocSlots
}