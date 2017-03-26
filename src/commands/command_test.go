package commands

import (
	"perror"
	"strings"
	"testing"
)

func TestCommand_Run(t *testing.T) {
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

	outPut := []struct {
		err error
		out string
	}{
		{nil, "Created a parking lot with 6 slots"},
		{nil, "Allocated slot number: 1"},
		{nil, "Allocated slot number: 2"},
		{nil, "Allocated slot number: 3"},
		{nil, "Allocated slot number: 4"},
		{nil, "Allocated slot number: 5"},
		{nil, "Allocated slot number: 6"},
		{nil, "Slot number 4 is free"},
		{nil, strings.Join(
			[]string{
				"Slot No.    Registration No     Colour",
				"------------------------------------------",
				"1           KA-01-HH-1234       White",
				"2           KA-01-HH-9999       White",
				"3           KA-01-BB-0001       Black",
				"5           KA-01-HH-2701       Blue",
				"6           KA-01-HH-3141       Black",
			}, "|"),
		},
		{nil, "Allocated slot number: 4"},
		{perror.ErrParkingFullCapacity, "Sorry, parking lot is full"},
		{nil, "KA-01-HH-1234, KA-01-HH-9999, KA-01-P-333"},
		{nil, "1, 2, 4"},
		{nil, "6"},
		{perror.ErrNotFound, "Not found"},
	}
	cmdMgr := NewManager()
	for idx, cmd := range cmdStrings {
		out, err := cmdMgr.Run(cmd)
		if cmd == "status" {
			for _, s := range strings.Split(outPut[idx].out, "|") {
				if !strings.Contains(out, s) {
					t.Error("Should have ", s)
				}
			}
		}
		if out != outPut[idx].out && cmd != "status" {
			t.Error("Should be expected output ", outPut[idx].out)
		}
		if err != outPut[idx].err {
			t.Error("Should be expected err ", outPut[idx].err)
		}
	}
}
