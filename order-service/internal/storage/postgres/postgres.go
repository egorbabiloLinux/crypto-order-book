package postgres

import (
	"context"
	"fmt"
	"time"

	"order-service/internal/domain/models/order"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/shopspring/decimal"
)

type Storage struct {
	db *pgxpool.Pool
}

type StorageWrapper struct {}

func NewStorage(url string) (*Storage, error) {
	const op = "storage.postgres.NewStorage"
	db, err := pgxpool.New(context.Background(), url)
	if err != nil {
		return nil, fmt.Errorf("%s: %v", op, err)
	}

	if err := db.Ping(context.Background()); err != nil {
		return nil, fmt.Errorf("%s: %v", op, err)
	}

	return &Storage{db: db}, nil
}

func NewStorageWrapper() *StorageWrapper {
	return &StorageWrapper{}
}

func (s *StorageWrapper) SaveOrder(
		userId 	  int64,
		price     decimal.Decimal,
		amount    decimal.Decimal,
		remaining decimal.Decimal,
		side      order.SideType,
		orderType order.OrderType,	
		status 	  order.StatusType,
		createdAt time.Time,
		updatedAt time.Time,
) (int64, error) {
	return 1, nil
}