// Copyright 2017 Krishna Kumar. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package commands with 'leave' command implementation
package commands

import (
	"fmt"
	"perror"
	"ptypes"
	"slot"
	"store"
	"strconv"
)

// CmdLeave defined arguments and related methods
type CmdLeave struct {
	Command
	SlotNumber ptypes.Index
}

// NewCmdLeave new leave command instance
func NewCmdLeave() *CmdLeave {
	var cmd = new(CmdLeave)
	cmd.Cmd = "leave"
	return cmd
}

// Help to print help of leave command
func (cl *CmdLeave) Help() {
	fmt.Println("No help found")
}

// Parse to parse arguments
func (cl *CmdLeave) Parse(argString string) error {
	cl.Command.Parse(argString)
	if perror.Empty != cl.Args[0] {
		val, err := strconv.ParseUint(cl.Args[0], 0, 64)
		if nil != err {
			return perror.ErrInvalidParams
		}
		cl.SlotNumber = ptypes.Index(val)
	}
	return nil
}

// Verify to check the provided parameters are valid or not
func (cl *CmdLeave) Verify() error {
	if !slot.IsValidSlotNumber(cl.SlotNumber) {
		return perror.ErrInvalidParams
	}
	return nil
}

// Run to execute the command and provide result
func (cl *CmdLeave) Run() (string, error) {
	pC := store.Get().GetParkingCenter()
	oSlot, err := pC.RemoveVehicleBySlotNumber(cl.SlotNumber)
	if nil == err {
		cl.OutPut = fmt.Sprintf(
			"Slot number %v is free",
			oSlot.GetNumber(),
		)
	}
	return cl.OutPut, err
}
