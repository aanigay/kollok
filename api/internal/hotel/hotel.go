package hotel

import (
	"context"
)

type Hotel struct {
	Id          int64   `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Facilities  string  `json:"facilities"`
	Address     string  `json:"address"`
	Price       float64 `json:"price"`
	Rating      float32 `json:"rating"`
}

type Repository interface {
	GetHotel(ctx context.Context, id int64) (*Hotel, error)
	GetAll(ctx context.Context) ([]Hotel, error)
}

type Service interface {
	GetHotel(c context.Context, id int64) (*Hotel, error)
	GetAll(c context.Context) ([]Hotel, error)
}
