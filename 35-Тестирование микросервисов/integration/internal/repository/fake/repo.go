package fake

import (
	"context"
	"errors"
	"github.com/jackc/pgx/v4/pgxpool"
	"integration_testing/internal/domain"
	"sync/atomic"
)

type Repo struct {
	m map[string]domain.Item
}

const defaultCapacity = 100

var counter atomic.Uint64

func NewRepo(db *pgxpool.Pool) *Repo {
	return &Repo{
		m: make(map[string]domain.Item, defaultCapacity),
	}
}

func (r *Repo) Save(ctx context.Context, item domain.Item) (uint64, error) {
	if _, ok := r.m[item.Name]; ok {
		return 0, errors.New("item already exists")
	}
	counter.Add(1)
	item.ID = counter.Load()
	r.m[item.Name] = item

	return item.ID, nil
}

func (r *Repo) Get(ctx context.Context, name string) (domain.Item, error) {
	item, ok := r.m[name]
	if !ok {
		return domain.Item{}, errors.New("item not found")
	}

	return item, nil
}

func (r *Repo) Update(ctx context.Context, item domain.Item) error {
	if _, ok := r.m[item.Name]; !ok {
		return errors.New("item not found")
	}
	oldItem, _ := r.m[item.Name]
	item.ID = oldItem.ID
	r.m[item.Name] = item

	return nil
}
