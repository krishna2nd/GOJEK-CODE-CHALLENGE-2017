package commands

import (
	"fmt"
	. "perror"
	"strings"
	"store"
)

type CmdGetSlotNumWithColour struct {
	Command
	Color string
}

func NewCmdGetSlotNumWithColour() *CmdGetSlotNumWithColour {
	var cmd *CmdGetSlotNumWithColour = new(CmdGetSlotNumWithColour)
	cmd.Cmd = "slot_numbers_for_cars_with_colour"
	return cmd
}

func (this *CmdGetSlotNumWithColour) Help() {
	fmt.Println("No help found")
}

func (this *CmdGetSlotNumWithColour) Parse(argString string) error {
	this.Command.Parse(argString);
	this.Color = this.Args[0];
	return nil
}

func (this *CmdGetSlotNumWithColour) Verify() error {
	if Empty == this.Color {
		return ErrInvalidParams
	}
	return nil
}

func (this *CmdGetSlotNumWithColour) Run() (error, string) {
	var outPutList = []string {}
	pC := store.Get().GetParkingCenter();
	err, slots := pC.ReportVehicleByColor(this.Color)
	if nil == err {
		for _, s := range slots {
			outPutList = append(
				outPutList,
				fmt.Sprintf("%v",  s.GetNumber()),
			)
		}
	} else {
		outPutList = []string { err.Error() }
	}
	this.OutPut = strings.Join(outPutList, Comma)
	return err, this.OutPut
}
