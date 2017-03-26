package store

import (
	"parking"
	"config"
)

type Store struct {
	parkingCenter *parking.ParkingCenter
}
var store *Store;

func New() *Store {
	if nil == store {
		store = new(Store).init()
	}
	return store
}

func NewStore() *Store {
	return New()
}

func Get() *Store {
	return New();
}


func(this *Store) init() *Store {
	this.parkingCenter = parking.New(
		config.Start,
		config.Capacity,
	);
	return this;
}

func(this *Store) GetParkingCenter() *parking.ParkingCenter {
	return this.parkingCenter
}

func(this *Store) SetParkingCenter(pC *parking.ParkingCenter) {
	this.parkingCenter = pC
}