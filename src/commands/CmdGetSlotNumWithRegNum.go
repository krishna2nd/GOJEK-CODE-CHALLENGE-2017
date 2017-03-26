// Copyright 2017 Krishna Kumar. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package commands with 'slot_number_for_registration_number'
// command implementation
package commands

import (
	"fmt"
	"perror"
	"store"
	"strings"
)

// CmdGetSlotNumWithRegNum defined arguments and related methods
type CmdGetSlotNumWithRegNum struct {
	Command
	RegistrationNumber string
}

// NewCmdGetSlotNumWithRegNum new command instance
func NewCmdGetSlotNumWithRegNum() *CmdGetSlotNumWithRegNum {
	var cmd = new(CmdGetSlotNumWithRegNum)
	cmd.Cmd = "slot_number_for_registration_number"
	return cmd
}

// Help to print command help information
func (cgs *CmdGetSlotNumWithRegNum) Help() {
	fmt.Println("No help found")
}

// Parse to parse arguments
func (cgs *CmdGetSlotNumWithRegNum) Parse(argString string) error {
	cgs.Command.Parse(argString)
	cgs.RegistrationNumber = cgs.Args[0]
	return nil
}

// Verify to check the provided parameters are valid or not
func (cgs *CmdGetSlotNumWithRegNum) Verify() error {
	if perror.Empty == cgs.RegistrationNumber {
		return perror.ErrInvalidParams
	}
	return nil
}

// Run to execute the command and provide result
func (cgs *CmdGetSlotNumWithRegNum) Run() (string, error) {
	var outPutList = []string{}
	pC := store.Get().GetParkingCenter()
	slots, err := pC.ReportVehicleByNumber(
		cgs.RegistrationNumber,
	)
	if nil == err {
		for _, s := range slots {
			outPutList = append(
				outPutList,
				fmt.Sprintf("%v", s.GetNumber()),
			)
		}
	} else {
		outPutList = []string{
			err.Error(),
		}
	}
	cgs.OutPut = strings.Join(outPutList, perror.Comma)
	return cgs.OutPut, err
}
