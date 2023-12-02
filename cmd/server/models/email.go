package models

import (
	"github.com/uptrace/bun"
	"time"
)

type Email struct {
	bun.BaseModel `bun:"table:emails,alias:e"`
	ID            int64     `bun:"id,pk,autoincrement" json:"id"`
	Account       string    `bun:"account,notnull" json:"accountName"`
	ToAddress     string    `bun:"to_address,notnull" json:"toAddress"`
	FromAddress   string    `bun:"from_address,notnull" json:"fromAddress"`
	Subject       string    `bun:"subject,notnull" json:"subject"`
	Body          string    `bun:"body,notnull" json:"body"`
	CreatedAt     time.Time `bun:"created_at,notnull" json:"createdAt"`
	UpdatedAt     time.Time `bun:"updated_at,notnull" json:"updatedAt"`
	SentAt        time.Time `bun:"sent_at" json:"sentAt"`
	Errors        string    `bun:"errors" json:"errors"`
}
