// Copyright 2017 Krishna Kumar. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package vehicle implements a simple vehicle object handling package.
// It defines a type, Vehicle,
// New method will be accepting vehicle number and Color properties
// GetColour(), GetNumber() Defined to access properties independently.
package vehicle

// Vehicle stores number and Color of vehicle
type Vehicle struct {
	// Number and Color is defined as string
	// eg
	//      Number:"KA-01-HH-1234"
	//      Color: "White"
	Number, Color string
}

// New - Package based New Object creation function
//  @params:
//      Number: string
//      Color: string
//  @return:
//		Vehicle: *Object
func New(number, color string) *Vehicle {
	return new(Vehicle).Init(number, color)
}

// Init - Initialise created object
//  @params:
//      Number: string
//      Color: string
//  @return:
//		Vehicle: *Object
func (v *Vehicle) Init(number, color string) *Vehicle {
	v.Color = color
	v.Number = number
	return v
}

// GetColour -  Get value of colour vehicle property
//  @params: (void)
//  @return:
//		Color: String
func (v *Vehicle) GetColour() string {
	return v.Color
}

// GetNumber - Get value of vehicle number property
//  @params: (void)
//  @return:
//		Number: String
func (v *Vehicle) GetNumber() string {
	return v.Number
}

// IsEquals - check object equality
//  @params: ({*Object} Vehicle)
//  @return:
//		Number: bool
func (v *Vehicle) IsEquals(vehicle *Vehicle) bool {
	return v.Number == vehicle.GetNumber() &&
		v.GetColour() == vehicle.GetColour()
}
