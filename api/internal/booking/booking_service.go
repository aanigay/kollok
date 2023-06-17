package booking

import (
	"context"
)

type service struct {
	Repository
}

func NewService(repository Repository) Service {
	return &service{
		repository,
	}
}

func (s *service) CreateBooking(c context.Context, req *CreateBookingReq) (*Booking, error) {
	dish, err := s.Repository.CreateBooking(c, req)
	if err != nil {
		return &Booking{}, err
	}
	return dish, nil
}
