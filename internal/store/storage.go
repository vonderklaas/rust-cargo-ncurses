package store

import (
	"context"
	"database/sql"
)

type Storage struct {
	Posts interface {
		Create(context.Context) error
	}
	Users interface {
		Create(context.Context) error
	}
}

func NewStorage(db *sql.DB) Storage {
	return Storage{
		Posts: &PostsStore{
			db: db,
		},
		Users: &UsersStore{
			db: db,
		},
	}
}
