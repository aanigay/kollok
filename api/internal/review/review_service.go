package review

import (
	"context"
	"github.com/golang-jwt/jwt/v4"
)

type service struct {
	Repository
}

func NewService(repository Repository) Service {
	return &service{
		repository,
	}
}

type Claims struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
	Role     string `json:"role"`
	jwt.RegisteredClaims
}

func (s *service) GetReview(c context.Context, hotelId int64) ([]Review, error) {
	d, err := s.Repository.GetReview(c, hotelId)
	if err != nil {
		return d, err
	}
	return d, nil
}

func (s *service) CreateTask(c context.Context, req *CreateReviewReq) (*Review, error) {
	dish, err := s.Repository.CreateReview(c, req)
	if err != nil {
		return &Review{}, err
	}
	return dish, nil
}
