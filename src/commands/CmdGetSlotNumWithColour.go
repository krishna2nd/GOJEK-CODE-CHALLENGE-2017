// Copyright 2017 Krishna Kumar. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package commands with 'slot_numbers_for_cars_with_colour'
// command implementation
package commands

import (
	"fmt"
	"perror"
	"store"
	"strings"
)

// CmdGetSlotNumWithColour defined arguments and related methods
type CmdGetSlotNumWithColour struct {
	Command
	Color string
}

// NewCmdGetSlotNumWithColour new command instance
func NewCmdGetSlotNumWithColour() *CmdGetSlotNumWithColour {
	var cmd = new(CmdGetSlotNumWithColour)
	cmd.Cmd = "slot_numbers_for_cars_with_colour"
	return cmd
}

// Help to print command help information
func (cgc *CmdGetSlotNumWithColour) Help() {
	fmt.Println("No help found")
}

// Parse to parse arguments
func (cgc *CmdGetSlotNumWithColour) Parse(argString string) error {
	cgc.Command.Parse(argString)
	cgc.Color = cgc.Args[0]
	return nil
}

// Verify to check the provided parameters are valid or not
func (cgc *CmdGetSlotNumWithColour) Verify() error {
	if perror.Empty == cgc.Color {
		return perror.ErrInvalidParams
	}
	return nil
}

// Run to execute the command and provide result
func (cgc *CmdGetSlotNumWithColour) Run() (string, error) {
	var outPutList = []string{}
	pC := store.Get().GetParkingCenter()
	slots, err := pC.ReportVehicleByColor(cgc.Color)
	if nil == err {
		for _, s := range slots {
			outPutList = append(
				outPutList,
				fmt.Sprintf("%v", s.GetNumber()),
			)
		}
	} else {
		outPutList = []string{err.Error()}
	}
	cgc.OutPut = strings.Join(outPutList, perror.Comma)
	return cgc.OutPut, err
}
