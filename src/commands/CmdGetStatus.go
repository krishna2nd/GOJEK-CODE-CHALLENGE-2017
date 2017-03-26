package commands

import (
	"fmt"
	"strings"
	"store"
	. "perror"
)

type CmdGetStatus struct {
	Command
}

func NewCmdGetStatus() *CmdGetStatus {
	var cmd *CmdGetStatus = new(CmdGetStatus)
	cmd.Cmd = "status"
	return cmd
}

func (this *CmdGetStatus) Help() {
	fmt.Println("No help found")
}

func (this *CmdGetStatus) Parse(argString string) error {
	this.Command.Parse(argString);
	return nil
}

func (this *CmdGetStatus) Verify() error {
	return nil
}

func (this *CmdGetStatus) Run() (error, string) {
	var outPutList = []string {
		fmt.Sprintf("%-12s%-20s%-10s",
			"Slot No.",
			"Registration No",
			"Colour",
		),
	}
	pC := store.Get().GetParkingCenter();
	err, slots := pC.ReportFilledSlots()
	if nil == err {
		for _, s := range slots {
			v := s.GetVehicle()
			outPutList = append(
				outPutList,
				fmt.Sprintf(
					"%-12v%-20v%-10v",
					s.GetNumber(),
					v.GetNumber(),
					v.GetColour(),
				),
			)
		}
		this.OutPut = strings.Join(outPutList, NewLine)
	} else {
		outPutList = append(
			outPutList,
			"No Data Found",
		)
	}
	return nil, this.OutPut
}
