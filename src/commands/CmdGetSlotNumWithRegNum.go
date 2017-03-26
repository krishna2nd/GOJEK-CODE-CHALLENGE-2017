package commands

import (
	"fmt"
	. "perror"
	"strings"
	"store"
)

type CmdGetSlotNumWithRegNum struct {
	Command
	RegistrationNumber string
}

func NewCmdGetSlotNumWithRegNum() *CmdGetSlotNumWithRegNum {
	var cmd *CmdGetSlotNumWithRegNum = new(CmdGetSlotNumWithRegNum)
	cmd.Cmd = "slot_number_for_registration_number"
	return cmd
}

func (this *CmdGetSlotNumWithRegNum) Help() {
	fmt.Println("No help found")
}

func (this *CmdGetSlotNumWithRegNum) Parse(argString string) error {
	this.Command.Parse(argString);
	this.RegistrationNumber = this.Args[0];
	return nil
}

func (this *CmdGetSlotNumWithRegNum) Verify() error {
	if Empty == this.RegistrationNumber {
		return ErrInvalidParams
	}
	return nil
}

func (this *CmdGetSlotNumWithRegNum) Run() (error, string) {
	var outPutList = []string {}
	pC := store.Get().GetParkingCenter();
	err, slots := pC.ReportVehicleByNumber(
		this.RegistrationNumber,
	)
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
