package booking

import (
	"context"
	"database/sql"
)

type DBTX interface {
	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
	PrepareContext(context.Context, string) (*sql.Stmt, error)
	QueryContext(context.Context, string, ...interface{}) (*sql.Rows, error)
	QueryRowContext(context.Context, string, ...interface{}) *sql.Row
}

type repository struct {
	db DBTX
}

func NewRepository(db DBTX) Repository {
	return &repository{db: db}
}

func (r *repository) CreateBooking(ctx context.Context, booking *CreateBookingReq) (*Booking, error) {
	b := Booking{
		HotelId:   booking.HotelId,
		Arrival:   booking.Arrival,
		Departure: booking.Departure,
	}
	var lastInsertId int
	query := "INSERT INTO \"booking\"(hotelId, arrival, departure) VALUES ($1, $2, $3) returning id"
	err := r.db.QueryRowContext(ctx, query, booking.HotelId, booking.Arrival, booking.Departure).Scan(&lastInsertId)
	if err != nil {
		return &Booking{}, err
	}

	b.Id = int64(lastInsertId)
	return &b, nil
}
