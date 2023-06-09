package postgres

import (
	"Impact/models"
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
)

type roomRepo struct {
	db *pgxpool.Pool
}

func NewRoomRepo(db *pgxpool.Pool) *roomRepo {
	return &roomRepo{
		db: db,
	}
}

func (r roomRepo) GetRoom(ctx context.Context, id int) (models.Room, error) {
	query := `SELECT * FROM room WHERE id=$1`
	var room models.Room
	err := r.db.QueryRow(ctx, query, id).Scan(&room.ID, &room.Name, &room.Type, &room.Capacity)
	if err != nil {
		return models.Room{}, err
	}
	return room, nil
}

func (r roomRepo) GetRooms(ctx context.Context, request models.RoomsRequest) (models.RoomsResponse, error) {
	query := `SELECT * FROM room`
	if request.Type != "" {
		query += ` WHERE type='` + request.Type + `'`
	}
	if request.Search != "" {
		if request.Type == "" {
			query += ` WHERE`

		} else {
			query += ` AND`
		}
		query += ` name LIKE '%` + request.Search + `%'`
	}
	query += ` LIMIT ` + fmt.Sprintf("%v", request.PageSize) + ` OFFSET ` + fmt.Sprintf("%v", request.PageSize*(request.Page-1)) + `;`

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return models.RoomsResponse{}, err
	}
	defer rows.Close()

	var rooms []models.Room
	var room models.Room
	for rows.Next() {
		err := rows.Scan(&room.ID, &room.Name, &room.Type, &room.Capacity)
		if err != nil {
			return models.RoomsResponse{}, err
		}
		rooms = append(rooms, room)
	}

	var count int
	err = r.db.QueryRow(ctx, `SELECT COUNT(*) FROM room;`).Scan(&count)
	if err != nil {
		return models.RoomsResponse{}, err
	}

	return models.RoomsResponse{
		Page:     request.Page,
		Count:    count,
		PageSize: request.PageSize,
		Results:  rooms,
	}, nil

}
