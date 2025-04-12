package models

import (
	"database/sql"
	"time"
)

type Category struct {
	ID         int            `json:"id"`
	Name       string         `json:"name"`
	CreatedAt  time.Time      `json:"created_at"`
	CreatedBy  sql.NullString `json:"created_by,omitempty"`
	ModifiedAt sql.NullTime   `json:"modified_at,omitempty"`
	ModifiedBy sql.NullString `json:"modified_by,omitempty"`
}
