package hotel

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
func (s *service) GetHotel(c context.Context, id int64) (*Hotel, error) {
	d, err := s.Repository.GetHotel(c, id)
	if err != nil {
		return &Hotel{}, err
	}
	return d, nil
}

func (s *service) GetAll(c context.Context) ([]Hotel, error) {
	dishes, err := s.Repository.GetAll(c)
	if err != nil {
		return dishes, err
	}
	return dishes, nil
}
