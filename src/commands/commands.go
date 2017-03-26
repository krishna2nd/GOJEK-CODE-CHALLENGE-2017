// Copyright 2017 Krishna Kumar. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package commands implements the basic shell command execution process
// Includes help, parse, verify, run
package commands

import (
	"perror"
	"strings"
)

// IManager should have behaviour run and a base parse
type IManager interface {
	Parse() error
	Run() (string, error)
}

// Manager handles requested command and available command's list
type Manager struct {
	cmd, argString string
	Commands       map[string]ICommand
}

// NewManager return command manager
func NewManager() *Manager {
	mgrCmd := new(Manager)
	mgrCmd.Commands = make(map[string]ICommand)

	mgrCmd.Register(NewCmdCreateParkingLot())
	mgrCmd.Register(NewCmdPark())
	mgrCmd.Register(NewCmdLeave())
	mgrCmd.Register(NewCmdGetStatus())
	mgrCmd.Register(NewCmdGetRegNumWithColour())
	mgrCmd.Register(NewCmdGetSlotNumWithColour())
	mgrCmd.Register(NewCmdGetSlotNumWithRegNum())
	return mgrCmd
}

// Register Command registration with manager
func (cm *Manager) Register(cmd ICommand) {
	cmdName := cmd.GetName()
	cm.Commands[cmdName] = cmd
}

// IsValidCommad verifies the requested command is valid or not
func (cm *Manager) IsValidCommad(cmdName string) bool {
	_, ok := cm.Commands[cmdName]
	return ok
}

// Parse requested command and arguments
func (cm *Manager) Parse(cmdString string) error {
	results := strings.SplitN(cmdString, perror.Space, 2)
	cm.cmd = results[0]
	if len(results) > 1 {
		cm.argString = results[1]
	}
	if perror.Empty == cm.cmd {
		return perror.ErrInvalidCommand
	}
	return nil
}

// Run the requested command and provide output
func (cm *Manager) Run(cmdString string) (string, error) {
	err := cm.Parse(cmdString)
	if nil != err {
		return perror.Empty, err
	}
	cmd, ok := cm.Commands[cm.cmd]
	if ok {
		cmd.Clear()
		err := cmd.Parse(cm.argString)
		if nil != err {
			return perror.Empty, perror.ErrCommandParsing
		}
		if nil == cmd.Verify() {
			return cmd.Run()
		}
		return perror.Empty, perror.ErrInvalidParams
	}
	return perror.Empty, perror.ErrInvalidCommand
}
