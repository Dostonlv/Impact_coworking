package postgres

import (
	"Impact/models"
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/lib/pq"
	_ "github.com/lib/pq"
	"log"
	"time"
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
	check, err := b.Check(ctx, request.Start, request.End)
	if err != nil {
		return models.BookingResponse{}, err
	}
	if check {
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
	return models.BookingResponse{}, nil
}

func (b bookingRoomRepo) Check(ctx context.Context, from, to time.Time) (bool, error) {
	startTime := from.Format("2006-01-02 15:04")
	endTime := to.Format("2006-01-02 15:04")
	fmt.Println(startTime, endTime)
	roomID := 2

	query := `
		SELECT id, resident, period
		FROM booking
		WHERE room_id = $1 AND period && tsrange($2, $3, '[)')
	`
	rows, err := b.db.Query(ctx, query, roomID, from, to)
	if err != nil {
		return false, err
	}
	defer rows.Close()
	for rows.Next() {
		var (
			id        int
			resident  string
			period    pq.NullTime
			errResult error
		)
		err := rows.Scan(&id, &resident, &period)
		if err != nil {
			log.Fatal(err)

		}
		errResult = fmt.Errorf("uzr, siz tanlagan vaqtda xona band")
		return false, errResult
	}
	if err = rows.Err(); err != nil {
		return false, err
	}
	return true, nil
}
