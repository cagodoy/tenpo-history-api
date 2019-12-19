package postgres

import (
	"github.com/Masterminds/squirrel"
	history "github.com/cagodoy/tenpo-history-api"
	"github.com/jmoiron/sqlx"
)

// HistoryStore ...
type HistoryStore struct {
	Store *sqlx.DB
}

// HistoryListByUserID ...
func (us *HistoryStore) HistoryListByUserID(q *history.Query) ([]*history.History, error) {
	query := squirrel.
		Select("*").
		From("history").
		Where(squirrel.Eq{"user_id": q.UserID}).
		Where("deleted_at is null")

	sql, args, err := query.PlaceholderFormat(squirrel.Dollar).ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := us.Store.Queryx(sql, args...)
	if err != nil {
		return nil, err
	}

	uu := make([]*history.History, 0)

	for rows.Next() {
		u := &history.History{}
		if err := rows.StructScan(u); err != nil {
			return nil, err
		}

		uu = append(uu, u)
	}

	return uu, nil
}

// HistoryCreate ...
func (us *HistoryStore) HistoryCreate(u *history.History) error {
	sql, args, err := squirrel.
		Insert("history").
		Columns("user_id", "latitude", "longitude").
		Values(u.UserID, u.Latitude, u.Longitude).
		Suffix("returning *").
		PlaceholderFormat(squirrel.Dollar).
		ToSql()

	if err != nil {
		return err
	}

	row := us.Store.QueryRowx(sql, args...)
	if err := row.StructScan(u); err != nil {
		return err
	}

	return nil
}

// HistoryList ...
func (us *HistoryStore) HistoryList() ([]*history.History, error) {
	query := squirrel.Select("*").From("history").Where("deleted_at is null")

	sql, args, err := query.PlaceholderFormat(squirrel.Dollar).ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := us.Store.Queryx(sql, args...)
	if err != nil {
		return nil, err
	}

	uu := make([]*history.History, 0)

	for rows.Next() {
		u := &history.History{}
		if err := rows.StructScan(u); err != nil {
			return nil, err
		}

		uu = append(uu, u)
	}

	return uu, nil
}
