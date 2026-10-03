package models

import (
	"time"
	"uuid"
)

type Post struct {
	ID        uuid.UUID `db:"id"`
	Title     string    `db:"title"`
	Category  string    `db:"category"`
	Author    string    `db:"author"`
	Content   string    `db:"content"`
	CreatedAt time.Time `db:"created_at"`
}

// updated at
