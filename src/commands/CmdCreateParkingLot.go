// Copyright 2017 Krishna Kumar. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package commands with 'create_parking_lot' command implementation
package commands

import (
	"config"
	"fmt"
	"parking"
	"perror"
	"ptypes"
	"store"
	"strconv"
)

// CmdCreateParkingLot defined arguments and related methods
type CmdCreateParkingLot struct {
	Command
	Capacity ptypes.Capacity
}

// NewCmdCreateParkingLot new command instance
func NewCmdCreateParkingLot() *CmdCreateParkingLot {
	var cmd = new(CmdCreateParkingLot)
	cmd.Cmd = "create_parking_lot"
	return cmd
}

// Help to print help of 'create_parking_lot' command
func (ccp *CmdCreateParkingLot) Help() string {
	return `🔸  create_parking_lot <slots count>
	Create parking lot slots.
	Eg: create_parking_lot 6`
}

// Parse to parse arguments
func (ccp *CmdCreateParkingLot) Parse(argString string) error {
	ccp.Command.Parse(argString)
	if perror.Empty != ccp.Args[0] {
		val, err := strconv.ParseUint(ccp.Args[0], 0, 64)
		if nil != err {
			return perror.ErrInvalidParams
		}
		ccp.Capacity = ptypes.Capacity(val)
	}
	return nil
}

// Verify to check the provided parameters are valid or not
func (ccp *CmdCreateParkingLot) Verify() error {
	if 1 > ccp.Capacity {
		return perror.ErrInvalidParams
	}
	return nil
}

// Run to execute the command and provide result
func (ccp *CmdCreateParkingLot) Run() (string, error) {
	pc := parking.New(config.Start,
		ccp.Capacity,
	)
	if nil != pc {
		store.Get().SetParkingCenter(pc)
		ccp.OutPut = fmt.Sprintf(
			"Created a parking lot with %v slots",
			ccp.Capacity,
		)
	} else {
		ccp.OutPut = perror.ErrCreateParkingCenter.Error()
	}
	return ccp.OutPut, nil
}
