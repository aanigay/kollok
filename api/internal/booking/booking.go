package booking

import (
	"context"
	"github.com/liamylian/jsontime"
	"time"
)

var json = jsontime.ConfigWithCustomTimeFormat

type Booking struct {
	Id        int64     `json:"id"`
	HotelId   int64     `json:"hotel_id"`
	Arrival   time.Time `json:"arrival" time_format:"sql_date" time_utc:"true"`
	Departure time.Time `json:"departure" time_format:"sql_date" time_utc:"true"`
}

type CreateBookingReq struct {
	HotelId   int64     `json:"hotel_id"`
	Arrival   time.Time `json:"arrival"`
	Departure time.Time `json:"departure"`
}

type Repository interface {
	CreateBooking(ctx context.Context, booking *CreateBookingReq) (*Booking, error)
}

type Service interface {
	CreateBooking(c context.Context, req *CreateBookingReq) (*Booking, error)
}
