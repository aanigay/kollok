package dish

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

func (s *service) GetDish(c context.Context, id int64) (*Dish, error) {
	d, err := s.Repository.GetDish(c, id)
	if err != nil {
		return &Dish{}, err
	}
	return d, nil
}

func (s *service) CreateDish(c context.Context, req *CreateDishReq) (*Dish, error) {
	dish, err := s.Repository.CreateDish(c, req)
	if err != nil {
		return &Dish{}, err
	}
	return dish, nil
}
func (s *service) UpdateDish(c context.Context, req *UpdateDishReq) (*Dish, error) {
	dish, err := s.Repository.UpdateDish(c, req)
	if err != nil {
		return &Dish{}, err
	}
	return dish, nil
}
func (s *service) DeleteDish(c context.Context, id int64) (*Dish, error) {
	d, err := s.Repository.GetDish(c, id)
	if err != nil {
		return &Dish{}, err
	}
	return d, nil
}

func (s *service) GetAll(c context.Context) ([]Dish, error) {
	dishes, err := s.Repository.GetAll(c)
	if err != nil {
		return dishes, err
	}
	return dishes, nil
}
