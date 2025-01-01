package expenses

import (
	"net/http"
	"strconv"

	"github.com/UAssylbek/Expense-Tracker-API/internal/db/expense"
	"github.com/UAssylbek/Expense-Tracker-API/pkg/httputils/request"
	"github.com/UAssylbek/Expense-Tracker-API/pkg/httputils/response"
)

type UpdateExpenseRequest struct {
	Data *expense.ModelExpense `json:"data"`
}

func (h *Expenses) UpdateExpense(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := h.logger.With("method", "UpdateExpense")

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

	// request parse
	requestBody := &UpdateExpenseRequest{}

	if err := request.JSON(w, r, requestBody); err != nil {
		log.ErrorContext(
			ctx,
			"failed to parse request body",
			"error", err,
		)
		http.Error(w, "failed to parse request body", http.StatusBadRequest)
		return
	}

	// db request
	if err := h.db.UpdateExpense(ctx, int64(id), requestBody.Data); err != nil {
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
		"success update expense",
		"id", id,
	)
}
