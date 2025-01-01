package expenses

import (
	"net/http"
	"strconv"

	"github.com/UAssylbek/Expense-Tracker-API/internal/db/expense"
	"github.com/UAssylbek/Expense-Tracker-API/pkg/httputils/response"
)

type FindExpenseResponse struct {
	Data *expense.ModelExpense `json:"data"`
}

func (h *Expenses) FindExpense(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := h.logger.With("method", "FindExpense")

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

	dbResp, err := h.db.FindExpense(ctx, int64(id))

	if err != nil {
		log.ErrorContext(
			ctx,
			"failed to query from db",
			"error", err,
		)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	resp := FindExpenseResponse{
		Data: dbResp,
	}

	if err := response.JSON(
		w,
		http.StatusOK,
		resp,
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
		"success find expense",
		"expense id", resp.Data.ID,
	)
}
