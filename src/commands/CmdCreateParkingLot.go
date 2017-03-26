package commands

import (
	"fmt"
	"store"
	"parking"
	"strconv"
	"config"
	. "ptypes"
	. "perror"
)

type CmdCreateParkingLot struct {
	Command
	Capacity Capacity
}

func NewCmdCreateParkingLot() *CmdCreateParkingLot {
	var cmd *CmdCreateParkingLot = new(CmdCreateParkingLot)
	cmd.Cmd = "create_parking_lot"
	return cmd
}

func (this *CmdCreateParkingLot) Help() {
	fmt.Println("No help found")
}

func (this *CmdCreateParkingLot) Parse(argString string) error {

	this.Command.Parse(argString);
	if Empty != this.Args[0] {
		val, err := strconv.ParseUint(this.Args[0], 0, 64)
		if nil != err {
			return ErrInvalidParams
		}
		this.Capacity = Capacity(val)
	}
	return nil
}

func (this *CmdCreateParkingLot) Verify() error {
	if (1 > this.Capacity) {
		return ErrInvalidParams
	}
	return nil
}

func (this *CmdCreateParkingLot) Run() (error, string) {
	store.Get().SetParkingCenter(
		parking.New(config.Start,
			this.Capacity,
		),
	)
	this.OutPut = fmt.Sprintf(
		"Created a parking lot with %v slots",
		this.Capacity,
	)
	return nil, this.OutPut
}
