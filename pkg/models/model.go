package models

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type BaseModel struct {
	ID        uuid.UUID `gorm:"primary_key;type:uuid"`
	CreatedAt time.Time
	UpdateAt  time.Time
	DeletedAt *sql.NullTime `gorm:"index"`
}
