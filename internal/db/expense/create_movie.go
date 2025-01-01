package expense

import (
	"context"
	"database/sql"
	"errors"
)

func (m *Expense) CreateExpense(ctx context.Context, insertData *ModelExpense) (*ModelExpense, error) {
	log := m.logger.With("method", "CreateExpense")

	stmt := `
INSERT INTO expense (cost, currency, title, description, category, posterUrl)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING id, cost, currency, title, description, category, posterUrl, created_at, updated_at
`

	row := m.db.QueryRowContext(ctx, stmt, insertData.Cost, insertData.Currency, insertData.Title, insertData.Description, insertData.Category, insertData.PosterURL)

	if err := row.Err(); err != nil {
		log.ErrorContext(ctx, "fail to insert to table expense", "error", err)
		return nil, err
	}

	expense := ModelExpense{}

	if err := row.Scan(
		&expense.ID,
		&expense.Cost,
		&expense.Currency,
		&expense.Title,
		&expense.Description,
		&expense.Category,
		&expense.PosterURL,
		&expense.CreatedAt,
		&expense.UpdatedAt,
	); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			log.ErrorContext(ctx, "no values was found", "error", err)
			return nil, nil
		}
		log.ErrorContext(ctx, "fail to scan expense", "error", err)
		return nil, err
	}

	log.InfoContext(ctx, "success insert to table expense")
	return &expense, nil
}
