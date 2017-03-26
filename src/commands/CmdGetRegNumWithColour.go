package commands

import (
	"fmt"
	. "perror"
	"strings"
	"store"
)

type CmdGetRegNumWithColour struct {
	Command
	Color string
}

func NewCmdGetRegNumWithColour() *CmdGetRegNumWithColour {
	var cmd *CmdGetRegNumWithColour = new(CmdGetRegNumWithColour)
	cmd.Cmd = "registration_numbers_for_cars_with_colour"
	return cmd
}

func (this *CmdGetRegNumWithColour) Help() {
	fmt.Println("No help found")
}

func (this *CmdGetRegNumWithColour) Parse(argString string) error {
	this.Command.Parse(argString);
	this.Color = this.Args[0];
	return nil
}

func (this *CmdGetRegNumWithColour) Verify() error {
	if Empty == this.Color {
		return ErrInvalidParams
	}
	return nil
}

func (this *CmdGetRegNumWithColour) Run() (error, string) {
	var outPutList = []string {}
	pC := store.Get().GetParkingCenter();
	err, slots := pC.ReportVehicleByColor(this.Color)
	if nil == err {
		for _, s := range slots {
			v := s.GetVehicle()
			outPutList = append(
				outPutList,
				fmt.Sprintf("%v", v.GetNumber()),
			)
		}
	} else {
		outPutList = []string { err.Error() }
	}
	this.OutPut = strings.Join(outPutList, Comma)
	return err, this.OutPut
}
