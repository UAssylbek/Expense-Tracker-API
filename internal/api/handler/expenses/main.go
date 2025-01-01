package expenses

import (
	"log/slog"

	"github.com/UAssylbek/Expense-Tracker-API/internal/db"
)

type Expenses struct {
	logger *slog.Logger
	db     *db.DB
}

func New(logger *slog.Logger, db *db.DB) *Expenses {
	return &Expenses{
		logger: logger,
		db:     db,
	}
}
