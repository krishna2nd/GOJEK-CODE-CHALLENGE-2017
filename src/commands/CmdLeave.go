package commands

import (
	"fmt"
	. "ptypes"
	. "perror"
	"slot"
	"store"
	"strconv"
)

type CmdLeave struct {
	Command
	SlotNumber Index
}

func NewCmdLeave() *CmdLeave {
	var cmd *CmdLeave = new(CmdLeave)
	cmd.Cmd = "leave"
	return cmd
}

func (this *CmdLeave) Help() {
	fmt.Println("No help found")
}

func (this *CmdLeave) Parse(argString string) error {
	this.Command.Parse(argString);
	if Empty != this.Args[0] {
		val, err := strconv.ParseUint(this.Args[0], 0, 64)
		if nil != err {
			return ErrInvalidParams
		}
		this.SlotNumber = Index(val)
	}
	return nil
}

func (this *CmdLeave) Verify() error {
	if !slot.IsValidSlotNumber(this.SlotNumber) {
		return ErrInvalidParams
	}
	return nil
}

func (this *CmdLeave) Run() (error, string) {
	pC := store.Get().GetParkingCenter();
	err, oSlot := pC.RemoveVehicleBySlotNumber(this.SlotNumber)
	if nil == err {
		this.OutPut = fmt.Sprintf(
			"Slot number %v is free",
			oSlot.GetNumber(),
		)
	}
	return err, this.OutPut
}
