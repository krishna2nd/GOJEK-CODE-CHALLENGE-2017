// Package commands test which test functionality test for
// Command file processing
package commands

import (
	"io/ioutil"
	"os"
	"testing"
)

var cmdStrings = []byte(
	`create_parking_lot 6
	park KA-01-HH-1234 White
	park KA-01-HH-9999 White
	park KA-01-BB-0001 Black
	park KA-01-HH-7777 Red
	park KA-01-HH-2701 Blue
	park KA-01-HH-3141 Black
	leave 4
	status
	park KA-01-P-333 White
	park DL-12-AA-9999 White
	registration_numbers_for_cars_with_colour White
	slot_numbers_for_cars_with_colour White
	slot_number_for_registration_number KA-01-HH-3141
	slot_number_for_registration_number MH-04-AY-1111`)

func TestFileCommandProcessor_Process(t *testing.T) {

	tmpfile, err := ioutil.TempFile("", "test.data")
	if err != nil {
		t.Fatal(err)
	}

	defer os.Remove(tmpfile.Name())

	if _, err := tmpfile.Write(cmdStrings); err != nil {
		t.Fatal(err)
	}
	if err := tmpfile.Close(); err != nil {
		t.Fatal(err)
	}

	err = NewFileCmdProcessor(tmpfile.Name()).Process()
	if nil != err {
		t.Error("Should not have error on process")
	}
}
