package storage

import (
	"Impact/models"
	"context"
)

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
	Check(ctx context.Context, from, to string) (bool, error)
}
