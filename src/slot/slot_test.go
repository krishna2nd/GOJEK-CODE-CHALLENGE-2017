package slot

import (
	"perror"
	"testing"
	"vehicle"
)

const (
	SlotInitNumberZero = "Created Unassigned slot must be zero"
	InvalidSlotNumber  = "Invalid slot number should show error"
	ValidSlotError     = "Valid slot number should not show error"
	VehicleSame        = "Vehicle should be same"
	AllocateUsingSlot  = "Should be error on allocate on using slot"
	ErrSlotInvalid     = "Expected 'ErrVehicleAssignInvalidSlot'"
	ErrSlotInvNum      = "Expected 'ErrSlotNumberInvalid'"
	ErrSlotAlloc       = "Expected 'ErrSlotAlreadyAllocated'"
	ErrFree            = "Slot Should be free"
)

var slots = []Slot{
	{0, vehicle.New("X", "W")},
	{1, vehicle.New("XX", "WW")},
	{2, vehicle.New("XXX", "WWW")},
}

func TestSlot_Init(t *testing.T) {
	var s *Slot = New()
	if s.Number != 0 {
		t.Error(SlotInitNumberZero)
	}
}

func TestSlot_SetNumber(t *testing.T) {
	var (
		s   *Slot
		err error
	)

	s = New()
	err, _ = s.SetNumber(0)
	if err == nil {
		t.Error(InvalidSlotNumber)
	}
	err, _ = s.SetNumber(SlotNumberLowerLimit)
	if err != nil {
		t.Error(ValidSlotError)
	}
}

func TestSlot_Allocate(t *testing.T) {
	var s *Slot

	for _, o := range slots {
		s = New()
		err, _ := s.Allocate(o.Vehicle)
		if err != perror.ErrVehicleAssignInvalidSlot {
			t.Error(ErrSlotInvalid)
		}

		err, _ = s.SetNumber(o.Number)
		if o.Number < SlotNumberLowerLimit {
			if err != perror.ErrSlotNumberInvalid {
				t.Error(ErrSlotInvNum)
			}
		}

		err, _ = s.Allocate(o.Vehicle)

		if o.Number < SlotNumberLowerLimit {
			if err != perror.ErrVehicleAssignInvalidSlot {
				t.Error(ErrSlotInvalid)
			}
		} else {
			if s.GetVehicle() != o.Vehicle {
				t.Error(VehicleSame)
			}
		}

		if !s.IsFree() {
			err, _ = s.Allocate(o.Vehicle)
			if err != perror.ErrSlotAlreadyAllocated {
				t.Error(ErrSlotAlloc)
			}
		}

		if !s.IsFree() {
			s.Free()
			if !s.IsFree() {
				t.Error(ErrFree)
			}
		}
	}
}
