// Copyright 2017 Krishna Kumar. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package commands base command features
// GetName, Clear, Parse, Verify, Run
package commands

import (
	"fmt"
	"perror"
	"strings"
)
// ICommand for base command's required behaviour
type ICommand interface {
	Help()
	GetName() string
	Parse(string) error
	Verify() error
	Run() (string, error)
	Clear()
}

// Command object structure
type Command struct {
	Cmd,
	InputArgs string
	OutPut string
	Args   []string
}

// NewCommand to create command instance
func NewCommand() *Command {
	var cmd = new(Command)
	return cmd
}

// Help to show usage
func (c *Command) Help() {
	fmt.Println("No help found")
}

// GetName to get the command name
func (c *Command) GetName() string {
	return c.Cmd
}

// Clear to clear the history data
func (c *Command) Clear() {
	c.InputArgs = perror.Empty
	c.Args = []string{}
	c.OutPut = perror.Empty
}

// Parse to help command to parse arguments from input string
func (c *Command) Parse(argString string) error {
	c.InputArgs = argString
	c.Args = strings.Split(argString, perror.Space)
	return nil
}

// Verify the provided Arguments
func (c *Command) Verify() error {
	return nil
}

// Run the command with arguments
func (c *Command) Run() (string, error) {
	return c.OutPut, nil
}
