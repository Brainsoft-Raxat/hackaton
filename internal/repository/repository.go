package repository

import (
	"context"
	"hackaton/internal/app/config"
	"hackaton/internal/app/conn"
)

type Repository struct {
	Postgres
}

type Postgres interface {
	DoSomething(ctx context.Context, name string) (value string, err error)
}

func New(conn conn.Conn, cfg *config.Config) *Repository {
	return &Repository{
		Postgres: NewPostgresRepository(conn.DB, cfg.Postgres),
	}
}
