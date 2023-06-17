package review

import (
	"context"
)

type Review struct {
	Id      int64   `json:"id"`
	HotelId int64   `json:"hotel_id"`
	Body    string  `json:"body"`
	Rating  float32 `json:"rating"`
}

type CreateReviewReq struct {
	HotelId int64   `json:"hotel_id"`
	Body    string  `json:"body"`
	Rating  float32 `json:"rating"`
}

type Repository interface {
	CreateReview(c context.Context, req *CreateReviewReq) (*Review, error)
	GetReview(ctx context.Context, hotelId int64) ([]Review, error)
}

type Service interface {
	CreateReview(c context.Context, req *CreateReviewReq) (*Review, error)
	GetReview(ctx context.Context, hotelId int64) ([]Review, error)
}
