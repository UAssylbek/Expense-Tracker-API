package expense

import (
	"context"
)

func (m *Expense) FindExpense(ctx context.Context, id int64) (*ModelExpense, error) {
	log := m.logger.With("method", "FindExpense")

	stmt := `
SELECT id, cost, currency, title, description, category, posterUrl, created_at, updated_at 
FROM expense
WHERE id = $1
`

	row := m.db.QueryRowContext(ctx, stmt, id)

	if err := row.Err(); err != nil {
		log.ErrorContext(ctx, "fail to query table expense", "error", err)
		return nil, err
	}

	expense := ModelExpense{}

	if err := row.Scan(
		&expense.ID,
		&expense.Title,
		&expense.Description,
		&expense.PosterURL,
		&expense.CreatedAt,
		&expense.UpdatedAt,
	); err != nil {
		log.ErrorContext(ctx, "fail to scan expense", "error", err)
		return nil, err
	}

	log.InfoContext(ctx, "success query table expense")
	return &expense, nil
}
