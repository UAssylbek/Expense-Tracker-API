package expenses

import (
	"fmt"
	"net/http"

	"github.com/UAssylbek/Expense-Tracker-API/internal/auth"
	"github.com/UAssylbek/Expense-Tracker-API/internal/db/expense"
	"github.com/UAssylbek/Expense-Tracker-API/pkg/httputils/request"
	"github.com/UAssylbek/Expense-Tracker-API/pkg/httputils/response"
)
const FTPRB = "failed to parse request body"
type CreateExpenseRequest struct {
	Data *expense.ModelExpense `json:"data"`
}

type CreateExpenseResponse struct {
	Data *expense.ModelExpense `json:"data"`
}

func (h *Expenses) CreateExpense(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := h.logger.With("method", "CreateExpense")

	user, ok := ctx.Value("user").(*auth.UserData)
	if !ok {
		log.ErrorContext(
			ctx,
			"failed to type cast user data",
		)
		http.Error(w, FTPRB, http.StatusBadRequest)
		return
	}

	fmt.Printf("user: %+v\n", *user)

	// request parse
	requestBody := &CreateExpenseRequest{}

	if err := request.JSON(w, r, requestBody); err != nil {
		log.ErrorContext(
			ctx,
			FTPRB,
			"error", err,
		)
		http.Error(w, FTPRB, http.StatusBadRequest)
		return
	}

	// db request
	dbResp, err := h.db.CreateExpense(ctx, requestBody.Data)

	if err != nil {
		log.ErrorContext(
			ctx,
			"failed to query from db",
			"error", err,
		)
		http.Error(w, "failed to query from db", http.StatusInternalServerError)
		return
	}

	if dbResp == nil {
		log.ErrorContext(
			ctx,
			"row is empty",
		)
		http.Error(w, "row is empty", http.StatusInternalServerError)
		return
	}

	// response
	resp := CreateExpenseResponse{
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
		"success insert expense",
		"expense id", resp.Data.ID,
	)
}
