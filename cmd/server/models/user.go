package models

import (
	"github.com/uptrace/bun"
	"time"
)

type User struct {
	bun.BaseModel `bun:"table:users,alias:u"`
	ID            int64     `bun:"id,pk,autoincrement" json:"id"`
	Username      string    `bun:"username,notnull" json:"username"`
	Password      string    `bun:"password,notnull" json:"password"`
	Email         string    `bun:"email,notnull" json:"email"`
	Role          string    `bun:"role,notnull" json:"role"`
	CreatedAt     time.Time `bun:"created_at,notnull" json:"createdAt"`
	UpdatedAt     time.Time `bun:"updated_at,notnull" json:"updatedAt"`
}
