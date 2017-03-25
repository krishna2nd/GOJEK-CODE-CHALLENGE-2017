// Copyright 2017 Krishna Kumar. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package vehicle implements a simple vehicle object handling package.
// It defines a type, Vehicle,
// New method will be accepting vehicle number and Color properties
// GetColour(), GetNumber() Defined to access properties independently.

package vehicle

type Vehicle struct {
	// Number and Color is defined as string
	// eg
	//      Number:"KA-01-HH-1234"
	//      Color: "White"
	Number, Color string
}

// Package based New Object creation function
//  @params:
//      Number: string
//      Color: string
//  @return:
//		Vehicle: *Object
func New(number, color string) *Vehicle {
	return new(Vehicle).Init(number, color)
}

// Initialise created object
//  @params:
//      Number: string
//      Color: string
//  @return:
//		Vehicle: *Object
func (this *Vehicle) Init(number, color string) *Vehicle {
	this.Color = color
	this.Number = number
	return this
}

// Get value of colour vehicle property
//  @params: (void)
//  @return:
//		Color: String
func (this *Vehicle) GetColour() string {
	return this.Color
}

// Get value of vehicle number property
//  @params: (void)
//  @return:
//		Number: String
func (this *Vehicle) GetNumber() string {
	return this.Number
}
