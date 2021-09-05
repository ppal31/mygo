package database

import (
	"context"
	"github.com/jmoiron/sqlx"
	"github.com/ppal31/mygo/internal/types"
)

type DataStore struct {
	ResourceStore ResourceStoreInterface
}

func New(db *sqlx.DB) *DataStore {
	// This is an important concept
	// We cannot do this : &DataStore{ResourceStore: ResourceStore{db: db}}
	// An interface value isn't the value of the concrete struct (as it has a variable size, this wouldn't be possible),
	// but it's a kind of pointer (to be more precise a pointer to the struct and a pointer to the type)
	return &DataStore{ResourceStore: &ResourceStore{db: db}}
}

type ResourceStoreInterface interface {
	Get(ctx context.Context, id int64) (*types.Resource, error)
	List(ctx context.Context, params types.Params) ([]*types.Resource, error)
}
