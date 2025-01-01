package expense

import (
	"context"
)

func (m *Expense) FindExpenses(ctx context.Context, offset, limit int) ([]ModelExpense, error) {
	log := m.logger.With("method", "FindExpenses")

	expenses := make([]ModelExpense, 0)

	stmt := `
SELECT id, title, description, posterUrl, created_at, updated_at 
FROM expense
OFFSET $1
LIMIT $2
`

	rows, err := m.db.QueryContext(ctx, stmt, offset, limit)
	if err != nil {
		log.ErrorContext(ctx, "fail to query table expense", "error", err)
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		expense := ModelExpense{}

		if err := rows.Scan(
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

		expenses = append(expenses, expense)
	}

	if err := rows.Err(); err != nil {
		log.ErrorContext(ctx, "fail to scan rows", "error", err)
		return nil, err
	}

	log.InfoContext(ctx, "success query table expense")
	return expenses, nil
}
