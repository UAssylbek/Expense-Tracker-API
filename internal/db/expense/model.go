package expense

import (
	"time"
)

type ModelExpense struct {
	ID          int        `json:"id"`
	Cost        int        `json:"cost"`
	Currency    string     `json:"currency"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Category    string     `json:"category"`
	PosterURL   string     `json:"poster_url"`
	CreatedAt   *time.Time `json:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at"`
}
