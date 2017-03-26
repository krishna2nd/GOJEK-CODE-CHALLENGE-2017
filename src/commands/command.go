/*
create_parking_lot
park
leave
status
registration_numbers_for_cars_with_colour
slot_numbers_for_cars_with_colour
slot_number_for_registration_number
*/

package commands

import (
	"fmt"
	"strings"
	. "perror"
)

type ICommand interface {
	Help()
	GetName() string
	Parse(string) error
	Verify() error
	Run() (error, string)
	Clear()
}

type Command struct {
	Cmd,
	InputArgs string
	OutPut     string
	Args []string
}

func NewCommand() *Command {
	var cmd *Command = new(Command)
	return cmd
}

func (this *Command) Help() {
	fmt.Println("No help found")
}

func (this *Command) GetName() string {
	return this.Cmd
}

func (this *Command) Clear() {
	this.InputArgs = Empty
	this.Args = []string{}
	this.OutPut = Empty
}

func (this *Command) Parse(argString string) error {
	this.InputArgs = argString
	this.Args  = strings.Split(argString, " ")
	return nil
}

func (this *Command) Verify() error {
	return nil
}

func (this *Command) Run() (error, string) {
	return nil, this.OutPut
}
