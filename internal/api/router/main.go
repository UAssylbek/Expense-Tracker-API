package router

import (
	"context"
	"net/http"

	"github.com/UAssylbek/Expense-Tracker-API/internal/api/handler"
	"github.com/UAssylbek/Expense-Tracker-API/internal/api/middleware"
)

type Router struct {
	router  *http.ServeMux
	handler *handler.Handler
	midd    *middleware.Middleware
}

func New(handler *handler.Handler, midd *middleware.Middleware) *Router {
	mux := http.NewServeMux()

	return &Router{
		router:  mux,
		handler: handler,
		midd:    midd,
	}
}

func (r *Router) Start(ctx context.Context) *http.ServeMux {
	r.auth(ctx)
	r.expenses(ctx)

	return r.router
}
