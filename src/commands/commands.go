package commands

import (
	"strings"
	 . "perror"

)

type ICommandManager interface {
	Run() (error, string)
}

type CommandManager struct {
	cmd, argString string
	Commands map[string]ICommand
}

func NewCommandManager() *CommandManager  {
	mgrCmd := new(CommandManager)
	mgrCmd.Commands = make(map[string]ICommand);
	
	mgrCmd.Register(NewCmdCreateParkingLot())
	mgrCmd.Register(NewCmdPark())
	mgrCmd.Register(NewCmdLeave())
	mgrCmd.Register(NewCmdGetStatus())
	mgrCmd.Register(NewCmdGetRegNumWithColour())
	mgrCmd.Register(NewCmdGetSlotNumWithColour())
	mgrCmd.Register(NewCmdGetSlotNumWithRegNum())
	return mgrCmd;
}

func (this *CommandManager) Register(cmd ICommand) {
	cmdName := cmd.GetName();
	this.Commands[cmdName] = cmd
}

func (this *CommandManager) IsValidCommad(cmdName string) bool {
	_, ok := this.Commands[cmdName]
	return ok
}
func (this *CommandManager) Parse(cmdString string) error {
	results := strings.SplitN(cmdString, Space, 2)
	this.cmd = results[0]
	if (len(results) > 1) {
		this.argString = results[1]
	}
	if Empty == this.cmd {
		return ErrInvalidCommand
	}
	return nil
}

func (this *CommandManager) Run(cmdString string) (error, string) {
	err := this.Parse(cmdString)
	if nil != err {
		return err, Empty
	}
	cmd, ok := this.Commands[this.cmd]
	if ok {
		cmd.Clear()
		err := cmd.Parse(this.argString)
		if nil != err {
			return ErrCommandParsing, Empty
		}
		if nil == cmd.Verify() {
			return cmd.Run()
		}
		return ErrInvalidParams, Empty
	}
	return ErrInvalidCommand, Empty
}


