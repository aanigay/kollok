package review

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

func (r *repository) CreateReview(ctx context.Context, reviewReq *CreateReviewReq) (*Review, error) {
	review := Review{
		HotelId: reviewReq.HotelId,
		Body:    reviewReq.Body,
		Rating:  reviewReq.Rating,
	}
	var lastInsertId int
	query := "INSERT INTO \"review\"(hotelId, body, rating) VALUES ($1, $2, $3) returning id"
	err := r.db.QueryRowContext(ctx, query, reviewReq.HotelId, reviewReq.Body, reviewReq.Rating).Scan(&lastInsertId)
	if err != nil {
		return &Review{}, err
	}

	review.Id = int64(lastInsertId)
	return &review, nil
}

func (r *repository) GetReview(ctx context.Context, hotelId int64) ([]Review, error) {
	reviews := make([]Review, 0)
	query := "SELECT id, hotelId, body, rating FROM \"review\" WHERE hotelId = $1"
	rows, err := r.db.QueryContext(ctx, query, hotelId)
	if err != nil {
		return reviews, nil
	}
	for rows.Next() {
		var review Review
		err := rows.Scan(&review.Id, &review.HotelId, &review.Body, &review.Rating)
		if err != nil {
			return reviews, err
		}
		reviews = append(reviews, review)
	}
	return reviews, nil
}
