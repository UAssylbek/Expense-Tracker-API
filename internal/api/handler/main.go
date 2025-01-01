package handler

import (
	"log/slog"

	"github.com/UAssylbek/Expense-Tracker-API/internal/api/handler/auth"
	"github.com/UAssylbek/Expense-Tracker-API/internal/api/handler/expenses"
	"github.com/UAssylbek/Expense-Tracker-API/internal/db"
)

type Handler struct {
	*auth.Auth
	*expenses.Expenses
}

func New(logger *slog.Logger, db *db.DB) *Handler {
	return &Handler{
		Auth:   auth.New(logger, db),
		Expenses: expenses.New(logger, db),
	}
}
