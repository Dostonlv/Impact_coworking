package storage

import (
	"Impact/models"
	"context"
)

// StorageI ...
type StorageI interface {
	Room() RoomRepoI
	CloseDB()
}

// RoomRepoI ...
type RoomRepoI interface {
	GetRoom(ctx context.Context, id int) (models.Room, error)
	GetRooms(ctx context.Context, request models.RoomsRequest) (models.RoomsResponse, error)
}
