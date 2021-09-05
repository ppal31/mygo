package database

import (
	"context"
	"github.com/jmoiron/sqlx"
	"github.com/ppal31/mygo/internal/types"
)

type ResourceStore struct {
	db *sqlx.DB
}

// Find finds the resource by id.
func (s *ResourceStore) Get(ctx context.Context, id int64) (*types.Resource, error) {
	dst := new(types.Resource)
	err := s.db.Get(dst, selectId, id)
	return dst, err
}

// Find finds the resource by id.
func (s *ResourceStore) List(ctx context.Context, params types.Params) ([]*types.Resource, error) {
	var dst []*types.Resource
	err := s.db.Select(&dst, selectWithParams, params.SizeOrDefault(), params.Offset())
	return dst, err
}

const resourceBase = `
SELECT
 id
,rtype
,name
,display_name
,url
FROM resources
`

const selectId = resourceBase + `WHERE id = $1`

const selectWithParams = resourceBase + `LIMIT $1 OFFSET $2`
