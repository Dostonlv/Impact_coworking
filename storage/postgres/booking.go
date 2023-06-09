package postgres

import (
	"Impact/models"
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
)

type bookingRoomRepo struct {
	db *pgxpool.Pool
}

func NewBookingRoomRepo(db *pgxpool.Pool) *bookingRoomRepo {
	return &bookingRoomRepo{
		db: db,
	}
}

func (b bookingRoomRepo) BookRoom(ctx context.Context, roomId int, request models.BookingRequest) (models.BookingResponse, error) {
	from := request.Start.Format("2006-01-02 15:04")
	to := request.End.Format("2006-01-02 15:04")
	fmt.Println(from, to)
	period := `[` + from + `, ` + to + `)`
	var boookingId int
	query := `INSERT INTO booking (room_id, resident, period) VALUES ($1, $2, $3) RETURNING id;`
	err := b.db.QueryRow(ctx, query, roomId, request.Resident.Name, period).Scan(&boookingId)
	if err != nil {
		return models.BookingResponse{}, err
	}

	return models.BookingResponse{Message: "xona muvaffaqiyatli band qilindi"}, nil
}
