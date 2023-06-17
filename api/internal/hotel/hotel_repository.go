package hotel

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

func (r *repository) GetAll(ctx context.Context) ([]Hotel, error) {
	hotels := make([]Hotel, 0)
	query := "SELECT * FROM hotel"
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return hotels, err
	}
	for rows.Next() {
		var hotel Hotel
		err = rows.Scan(&hotel.Id, &hotel.Name, &hotel.Description, &hotel.Facilities, &hotel.Address, &hotel.Price, &hotel.Rating)
		if err != nil {
			return hotels, err
		}
		hotels = append(hotels, hotel)
	}
	return hotels, nil
}

func (r *repository) GetHotel(ctx context.Context, id int64) (*Hotel, error) {
	hotel := Hotel{}
	query := "SELECT id, name, description, facilities, address, price, rating FROM \"hotel\" WHERE id = $1"
	err := r.db.QueryRowContext(ctx, query, id).Scan(&hotel.Id, &hotel.Name, &hotel.Description, &hotel.Facilities, &hotel.Address, &hotel.Price, &hotel.Rating)
	if err != nil {
		return &Hotel{}, nil
	}

	return &hotel, nil
}
