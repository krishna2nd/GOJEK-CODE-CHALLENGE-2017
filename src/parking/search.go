package parking

import "slot"

func (this *ParkingCenter) ReportVehicleByColor(number string) []*slot.Slot {
	_, oSlots := this.GetSlotsBy("color", number)
	return oSlots
}

func (this *ParkingCenter) ReportVehicleByNumber(number string) []*slot.Slot {
	_, oSlots := this.GetSlotsBy("number", number)
	return oSlots
}

func (this *ParkingCenter) ReportFreeSlots() []*slot.Slot {
	return this.getAllFreeSlot()
}
