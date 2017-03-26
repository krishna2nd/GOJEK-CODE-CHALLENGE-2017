package main

import (
	"commands"
	//"parking"
	//"fmt"
	//"config"
	"fmt"
)

func main() {
	cmdStrings := []string{
		"create_parking_lot 6",
		"park KA-01-HH-1234 White",
		"park KA-01-HH-9999 White",
		"park KA-01-BB-0001 Black",
		"park KA-01-HH-7777 Red",
		"park KA-01-HH-2701 Blue",
		"park KA-01-HH-3141 Black",
		"leave 4",
		"status",
		"park KA-01-P-333 White",
		"park DL-12-AA-9999 White",
		"registration_numbers_for_cars_with_colour White",
		"slot_numbers_for_cars_with_colour White",
		"slot_number_for_registration_number KA-01-HH-3141",
		"slot_number_for_registration_number MH-04-AY-1111",
	}
	cmdMgr := commands.NewManager()
	for _, cmd := range cmdStrings {
		err, out := cmdMgr.Run(cmd)
		fmt.Println(out, err)
	}
}

/*

mgr := New()

	err, out := mgr.Run("park test")
	fmt.Println(err, out)

type ParkingSlotManager struct {
	parkingCenter *parking.ParkingCenter
	cmdMgr *commands.CommandManager
}

func New() *ParkingSlotManager {
	mgr := new(ParkingSlotManager);
	mgr.cmdMgr = commands.NewCommandManager()
	mgr.parkingCenter = parking.New(config.Start, config.Capacity)
	return mgr;
}

func (this *ParkingSlotManager) Run(cmdString string) (error, string)  {
	return this.cmdMgr.Run(cmdString)
}
*/
