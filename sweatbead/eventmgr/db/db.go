package db

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ctxKey int

const (
	dbKey          ctxKey = 0
	defaultTimeout        = 1 * time.Second
)

type Storer interface {
	CreateSweat(ctx context.Context, sweat *Sweat) (sweatInfo Sweat, err error)
	ListSweats(ctx context.Context) (sweats []Sweat, err error)
	FindSweatByID(ctx context.Context, sweatID primitive.ObjectID) (sweat Sweat, err error)
	DeleteSweatByID(ctx context.Context, sweatID primitive.ObjectID) (err error)
}

type store struct {
	db *mongo.Database
}

func WithTimeout(ctx context.Context, timeout time.Duration, op func(ctx context.Context) error) (err error) {
	ctxWithTimeout, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	return op(ctxWithTimeout)
}

func WithDefaultTimeout(ctx context.Context, op func(ctx context.Context) error) (err error) {
	return WithTimeout(ctx, defaultTimeout, op)
}

func NewStorer(d *mongo.Database) Storer {
	return &store{
		db: d,
	}
}
