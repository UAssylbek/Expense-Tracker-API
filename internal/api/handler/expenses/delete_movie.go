package expenses

import (
	"net/http"
	"strconv"

	"github.com/UAssylbek/Expense-Tracker-API/pkg/httputils/response"
)

func (h *Expenses) DeleteExpense(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := h.logger.With("method", "DeleteExpense")

	idStr := r.PathValue("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.ErrorContext(
			ctx,
			"failed to convert id to int",
			"error", err,
		)
		http.Error(w, "failed to convert id to int", http.StatusBadRequest)
		return
	}

	// db request
	if err := h.db.DeleteExpense(ctx, int64(id)); err != nil {
		log.ErrorContext(
			ctx,
			"failed to query from db",
			"error", err,
		)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := response.JSON(
		w,
		http.StatusNoContent,
		nil,
	); err != nil {
		log.ErrorContext(
			ctx,
			"fail json",
			"error", err,
		)
		return
	}

	log.InfoContext(
		ctx,
		"success delete expense",
		"id", id,
	)
}
