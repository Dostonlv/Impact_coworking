package storage

import (
	"Impact/models"
	"context"
)

// for review

// StorageI ...
type StorageI interface {
	Room() RoomRepoI
	Booking() BookingRepoI
	CloseDB()
}

// RoomRepoI ...
type RoomRepoI interface {
	GetRoom(ctx context.Context, id int) (models.Room, error)
	GetRooms(ctx context.Context, request models.RoomsRequest) (models.RoomsResponse, error)
}

// BookingRepoI ...
type BookingRepoI interface {
	BookRoom(ctx context.Context, roomId int, request models.BookingRequest) (models.BookingResponse, error)
	Check(ctx context.Context, check models.Check) (bool, error)
	// get booking  rooms by roomId
	GetBookingRooms(ctx context.Context, roomId int) ([]models.GetBookingRoomsResponse, error)
	// get Booking rooms time and parse data and time then return available rooms
	//GetAvailableRooms(ctx context.Context, request models.AvailableRooms) ([]models.AvailableRooms, error)
}
