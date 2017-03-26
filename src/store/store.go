// Copyright 2017 Krishna Kumar. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package store handling center for Parking center instance
package store

import (
	"parking"
	"config"
)
// Store keeps a global reference for Parking Center
type Store struct {
	parkingCenter *parking.Center
}
// store : local storage
var store *Store;

// New : New store instance can be created
//  Singleton instance creation and return
//  @params: (void)
//  @return:
//		Store: *Object
func New() *Store {
	if nil == store {
		store = new(Store).init()
	}
	return store
}

// NewStore : store instance can be created
//  Singleton instance creation and return
//  @params: (void)
//  @return:
//		Store: *Object
func NewStore() *Store {
	return New()
}

// Get : store instance can be created
//  Singleton instance creation and return
//  @params: (void)
//  @return:
//		Store: *Object
func Get() *Store {
	return New();
}

// init : initialise created object
//  @params: (void)
//  @return:
//		Store: *Object
func(s *Store) init() *Store {
	s.parkingCenter = parking.New(
		config.Start,
		config.Capacity,
	);
	return s;
}

// GetParkingCenter : Return the current parking center instance
//  @params: (void)
//  @return:
//		ParkingCenter: *Object
func(s *Store) GetParkingCenter() *parking.Center {
	return s.parkingCenter
}


// SetParkingCenter : Set new ParkingCenter object in memory
//  @params: (void)
//  @return:
//		ParkingCenter: *Object
func(s *Store) SetParkingCenter(pC *parking.Center) {
	s.parkingCenter = pC
}
