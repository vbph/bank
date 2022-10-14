// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0

package db

import (
	"time"
)

type Account struct {
	ID                int64     `json:"id"`
	FullName          string    `json:"full_name"`
	Password          string    `json:"password"`
	Balance           int64     `json:"balance"`
	PasswordChangedAt time.Time `json:"password_changed_at"`
	CreatedAt         time.Time `json:"created_at"`
}

type Entry struct {
	ID        int64     `json:"id"`
	AccountID int64     `json:"account_id"`
	Amount    int64     `json:"amount"`
	CreatedAt time.Time `json:"created_at"`
}

type Transfer struct {
	ID        int64     `json:"id"`
	FromID    int64     `json:"from_id"`
	ToID      int64     `json:"to_id"`
	Amount    int64     `json:"amount"`
	CreatedAt time.Time `json:"created_at"`
}
