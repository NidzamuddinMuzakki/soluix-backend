package storage

import (
	"context"
	"time"
)

type PostStorer interface {
	Insert(ctx context.Context, post Post) error
	Update(ctx context.Context, post Post) error
	Delete(ctx context.Context, id string) error
	FindOne(ctx context.Context, id string) (Post, error)
}

type Post struct {
	ID        string     `json:"id"`
	Nama      string     `json:"nama"`
	Kategori  string     `json:"kategori"`
	Stok      int        `json:"stok"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
}
