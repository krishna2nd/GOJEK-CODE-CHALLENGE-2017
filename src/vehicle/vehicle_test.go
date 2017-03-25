// Simple unit testing on Vehicle object methods.

package vehicle

import (
	"testing"
)

var vehicles = []Vehicle{
	{"", ""},
	{"XXX", "CCC"},
}

func TestVehicle_GetColour(t *testing.T) {
	var v *Vehicle
	for _, o := range vehicles {
		v = New(o.Number, o.Color)
		if v.GetColour() != o.Color {
			t.Error("Colour should be same")
		}
	}
}

func TestVehicle_GetNumber(t *testing.T) {
	var v *Vehicle
	for _, o := range vehicles {
		v = New(o.Number, o.Color)
		if v.GetNumber() != o.Number {
			t.Error("Number should be same")
		}
	}
}
