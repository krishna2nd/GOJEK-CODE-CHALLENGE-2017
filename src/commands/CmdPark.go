package commands

import (
	"fmt"
	. "perror"
	"store"
	"vehicle"
)

type CmdPark struct {
	Command
	Vehicle *vehicle.Vehicle
}

func NewCmdPark() *CmdPark {
	var cmd *CmdPark = new(CmdPark)
	cmd.Cmd = "park"
	return cmd
}

func (this *CmdPark) Help() {
	fmt.Println("No help found")
}

func (this *CmdPark) Parse(argString string) error {
	this.Command.Parse(argString);
	return nil
}

func (this *CmdPark) Verify() error {
	if(2 != len(this.Args)) {
		return ErrInvalidParams
	}
	
	this.Vehicle = vehicle.New(this.Args[0], this.Args[1])
	return nil
}

func (this *CmdPark) Run() (error, string) {
	pC := store.Get().GetParkingCenter();
	err, oSlot := pC.AddVehicle(this.Vehicle)
	if nil == err {
		this.OutPut = fmt.Sprintf(
			"Allocated slot number: %v",
			oSlot.GetNumber(),
		)
	}
	return err, this.OutPut
}
