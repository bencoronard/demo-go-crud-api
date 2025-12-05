package resource

import (
	"context"
	"database/sql"

	"github.com/bencoronard/demo-go-crud-api/pkg"
)

type ResourceRepository interface {
	FindAll(ctx context.Context, p pkg.Pageable) (*pkg.Slice[*Resource], error)
	FindById(ctx context.Context, id int64) (*Resource, error)
	Save(ctx context.Context, ent *Resource) (*Resource, error)
	Delete(ctx context.Context, ent *Resource) error
}

type ResourceRepositoryImpl struct {
	db *sql.DB
}

func NewResourceRepositoryImpl(db *sql.DB) ResourceRepository {
	return &ResourceRepositoryImpl{db: db}
}

func (r *ResourceRepositoryImpl) FindAll(ctx context.Context, p pkg.Pageable) (*pkg.Slice[*Resource], error) {
	var s pkg.Slice[*Resource]
	return &s, nil
}

func (r *ResourceRepositoryImpl) FindById(ctx context.Context, id int64) (*Resource, error) {
	return nil, nil
}

func (r *ResourceRepositoryImpl) Save(ctx context.Context, ent *Resource) (*Resource, error) {
	return nil, nil
}

func (r *ResourceRepositoryImpl) Delete(ctx context.Context, ent *Resource) error {
	return nil
}
