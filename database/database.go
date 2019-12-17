package database

import (
	history "github.com/cagodoy/tenpo-history-api"
	"github.com/cagodoy/tenpo-history-api/database/postgres"
	"github.com/jmoiron/sqlx"
)

// Store ...
type Store interface {
	HistoryListByUserID(*history.Query) ([]*history.History, error)
	HistoryCreate(*history.History) error
	HistoryList() ([]*history.History, error)
}

// NewPostgres ...
func NewPostgres(dsn string) (Store, error) {
	db, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		return nil, err
	}

	return &postgres.HistoryStore{
		Store: db,
	}, nil
}
