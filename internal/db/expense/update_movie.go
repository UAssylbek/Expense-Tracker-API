package expense

import (
	"context"
	"fmt"
)

func (m *Expense) UpdateExpense(ctx context.Context, id int64, insertData *ModelExpense) error {
	log := m.logger.With("method", "UpdateExpense", "id", id)

	stmt := `
UPDATE expense
SET cost = $2, currency = $3, title = $4, description = $5, category = $6, posterUrl = $7
WHERE id = $1
`

	res, err := m.db.ExecContext(ctx, stmt, id, insertData.Title, insertData.Description, insertData.PosterURL)
	if err != nil {
		log.ErrorContext(ctx, "fail to update the table expense", "error", err)
		return err
	}

	num, err := res.RowsAffected()
	if err != nil {
		log.ErrorContext(ctx, "fail to update from the table expense", "error", err)
		return err
	}

	if num == 0 {
		log.WarnContext(ctx, "expense with id was not found", "id", id)
		return fmt.Errorf("expense with id was not found")
	}

	log.InfoContext(ctx, "success update the table expense")
	return nil
}
