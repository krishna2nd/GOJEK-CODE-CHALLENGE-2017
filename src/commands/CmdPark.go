// Copyright 2017 Krishna Kumar. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package commands with 'park' command implementation
package commands

import (
	"fmt"
	"perror"
	"store"
	"vehicle"
)

// CmdPark defined arguments and related methods
type CmdPark struct {
	Command
	Vehicle *vehicle.Vehicle
}

// NewCmdPark new park command instance
func NewCmdPark() *CmdPark {
	var cmd = new(CmdPark)
	cmd.Cmd = "park"
	return cmd
}

// Help to print help of park command
func (cp *CmdPark) Help() {
	fmt.Println("No help found")
}

// Parse to parse arguments
func (cp *CmdPark) Parse(argString string) error {
	cp.Command.Parse(argString)
	return nil
}

// Verify to check the provided parameters are valid or not
func (cp *CmdPark) Verify() error {
	if 2 != len(cp.Args) {
		return perror.ErrInvalidParams
	}

	cp.Vehicle = vehicle.New(cp.Args[0], cp.Args[1])
	return nil
}

// Run to execute the command and provide result
func (cp *CmdPark) Run() (string, error) {
	pC := store.Get().GetParkingCenter()
	oSlot, err := pC.AddVehicle(cp.Vehicle)
	if nil == err {
		cp.OutPut = fmt.Sprintf(
			"Allocated slot number: %v",
			oSlot.GetNumber(),
		)
	} else {
		cp.OutPut = err.Error()
	}
	return cp.OutPut, err
}
