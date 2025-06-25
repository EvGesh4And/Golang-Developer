package postgres

import (
	"context"
	"errors"
	"fmt"
	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"integration_testing/internal/domain"
	"time"
)

const (
	itemsTable = "items"
)

type Repo struct {
	db *pgxpool.Pool
}

func NewRepo(db *pgxpool.Pool) *Repo {
	return &Repo{
		db: db,
	}
}

func (r *Repo) Save(ctx context.Context, item domain.Item) (uint64, error) {
	tx, err := r.db.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return 0, fmt.Errorf("can't create tx: %w", err)
	}

	defer func() {
		if err != nil {
			rollbackErr := tx.Rollback(ctx)
			if rollbackErr != nil {
				err = errors.Join(err, rollbackErr)
			}
		}
	}()

	query, args, err := sq.
		Insert(itemsTable).
		Columns("name", "description", "created_at", "updated_at").
		Values(
			item.Name,
			item.Description,
			time.Now().Format(time.RFC3339),
			time.Now().Format(time.RFC3339),
		).
		Suffix("RETURNING id").
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		return 0, fmt.Errorf("can't build sql: %w", err)
	}

	rows, err := tx.Query(ctx, query, args...)
	if err != nil {
		return 0, fmt.Errorf("tx err: %w", err)
	}
	defer rows.Close()

	var itemID uint64
	for rows.Next() {
		if scanErr := rows.Scan(&itemID); scanErr != nil {
			return 0, fmt.Errorf("can't scan itemID: %w", scanErr)
		}
	}

	if err := tx.Commit(ctx); err != nil {
		return 0, fmt.Errorf("can't commit tx: %w", err)
	}

	return itemID, nil
}

func (r *Repo) Get(ctx context.Context, name string) (domain.Item, error) {

	// build
	query, args, err := sq.
		Select("id", "name", "description", "created_at", "updated_at").
		From(itemsTable).
		Where(sq.Eq{"name": name}).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		return domain.Item{}, fmt.Errorf("can't build query: %w", err)
	}

	// get
	item := domain.Item{}
	err = r.db.QueryRow(ctx, query, args...).Scan(&item.ID, &item.Name, &item.Description, &item.CreatedAt, &item.UpdatedAt)
	if err != nil {
		return domain.Item{}, fmt.Errorf("can't select orders: %w", err)
	}

	//for rows.Next() {
	//	scanErr := rows.Scan()
	//	if scanErr != nil {
	//		return domain.Item{}, fmt.Errorf("can't scan order: %w", scanErr)
	//	}
	//}

	return item, nil
}
