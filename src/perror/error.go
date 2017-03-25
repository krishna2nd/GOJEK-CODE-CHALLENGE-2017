package perror

import "errors"

var (
	ErrSlotNumberInvalid        = errors.New("slot: Please provide valid slot number 1 or greater.")
	ErrSlotNumberNotAssigned    = errors.New("slot: Slot number not assigned.")
	ErrSlotAlreadyAllocated     = errors.New("slot: Slots already allocated.")
	ErrSlotDuplicateNumber      = errors.New("slot: Please use unique number for slots.")
	ErrVehicleData              = errors.New("vehicle:Please enter a valid vehicle Number & Colour.")
	ErrVehicleDataAlreadyAdded  = errors.New("vehicle: Data already added.")
	ErrVehicleAssignInvalidSlot = errors.New("vehicle: Cannot allocate vehicle in invalid slot.")
)
